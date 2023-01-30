# Chain4Energy minter module - cfeminter

## Abstract

Chain4Energy minter module provides functionality of controlled token emissions. Tokens are emitted according to configuration.

## Contents

1. **[Concept](#concepts)**
2. **[Params](#parameters)**
3. **[State](#state)**
4. **[Events](#events)**
5. **[Queries](#queries)**
6. **[Genesis validations](#genesis-validations)**

## Concepts

The purpose of `cfeminter` module is to provide token emission mechanism.

### Token emission mechanism

Token emission mechanism mints calculated amount of tokens per each block. 
Token amount is calculated accordingly to cfeminter module configuration params.
Tokens minting process is divided into separate minters where each minter has
different minting rules. Those minting rules are defined within cfeminter 
module configuration params.
Simply, minting process configuration is a list of ordered minters,
where each minter has its own start and end time. Last minter cannot have end time because it
must be defined to work infinitely.

### Minters 

Ordered list of minters defines whole token emission process.
End time of one minter is a start time of the next minter in the minters list.
Each minter has its own type assigned.

### Minting type

Minting type defines general rules of token emission. Each minter has its
specific minting type. Currently, cfeminter module supports following minting types:
- no minting
- linear minting
- exponential step minting

Each minting type has its own specific set of parameters modifying token emission.

#### No minting

No minting is a simple minting type that mints nothing.
This minting type has no parameters.

#### Linear minting

Linear minting is block time based minting type. It mints configured amount of 
tokens within minter linearly. This minting type requires minter with end time 
since given amount of token needs to be minted in finite time period. So this 
minting type cannot be configured as a type of last period.

Minter type parameters:
* amount - amount of tokens to mint within minter period

#### Exponential step minting

Exponential step minting is block time based minting type. It mints configured amount
of tokens within minter, where it divides this minter into smaller subminters of 
equal length. Then within each subminter expected amount is minted, linearly. Expected 
amount of subminter minted tokens is equal to tokens minted by previous subminter 
multiplied by configured factor. For example initial minter amount is 40 million, 
multiplying factor set to 0.5 and step duration is four years, then:
* 1st subminter (first 4 years) mints 40 millions linearly 
* 2nd subminter (second 4 years) mints 20 millions linearly
* 3rd subminter (third 4 years) mints 10 millions linearly
* 4th subminter (fourth 4 years) mints 5 millions linearly
and so on.

This minter can mint infinitely.

Minter type parameters: 
* step duration - period of time mint amount is emitted
* amount - amount to mint for the first period
* amount multiplier - amount multiplying factor

## Examples

### Four years halving minting that starts with 40 million tokens and step duration set at 4 years

Minter configuration:
* minting start: now
* Amount of minter Minters: 1
* Minter 1:
    * end time: null
    * type: exponential step minting
    * exponential step minting parameters:
        * step duration: 4 years
        * amount: 40 millions
        * amount multiplier: 0.5

Result:
* first 4 years mints 40 millions 
* second 4 years mints 20 millions
* third 4 years mints 10 millions
and so on

### Linear minting of 100 million of token during period of 10 years, next no emission

Minter configuration:

* minting start: now
* Amount of minter Minters: 2
* Minter 1:
    * end time: 10 years from now
    * type: linear minting
    * linear minting parameters:
        * amount: 100 millions
* Minter 2:
    * end time: null
    * type: no minting

Result:
* 10 millions yearly for 10 years

## Parameters

The Chain4Energy minter module contains the following configurations parameters:

| Key           | Type         | Description                  |
|---------------|--------------|------------------------------|
| mint_denom    | string       | Denom of minting token       |
| minter_config | MinterConfig | Token emission configuration |

### MinterConfig type

| Param      | Type            | Description               |
|------------|-----------------|---------------------------|
| start_time | Time            | Token emission start time |
| minters    | List of Minters | list of minters           |

### Minter type

| Param                    | Type                   | Description                                                                                             |
|--------------------------|------------------------|---------------------------------------------------------------------------------------------------------|
| sequence_id              | uint32                 | Minter ordering id                                                                                      |
| end_time                 | Time                   | Minter end time                                                                                         |
| type                     | Enum string            | Minter period type. Allowed values:<br>- NO_MINTING<br>- LINEAR_MINTING <br>- EXPONENTIAL_STEP_MINTING; |
| linear_minting           | LinearMinting          | Linear minting configuration                                                                            |
| exponential_step_minting | ExponentialStepMinting | Exponential step minting configuration                                                                  |

### LinearMinting type

| Param    | Type          | Description                                  |
|----------|---------------|----------------------------------------------|
| amount   | sdk.Int       | An amount to mint linearly during the period |

### ExponentialStepMinting type

| Param             | Type      | Description                          |
|-------------------|-----------|--------------------------------------|
| step_duration     | uint32    | period of time of token emission     |
| amount            | sdk.Int   | amount to mint during "stepDuration" |
| amount_multiplier | sdk.Dec   | amount multiplying factor            |

### Example params

1. Four years halving minting that starts with 40 million tokens and step duration set at 4 years

```json
{
  "params": {
    "mint_denom": "uc4e",
    "minter_config": {
      "start_time": "2022-07-05T00:00:00Z",
      "minters": [
        {
          "sequenceId": 1,
          "end_time": null,
          "type": "EXPONENTIAL_STEP_MINTING",
          "linear_minting": null,
          "exponential_step_minting": {
            "step_duration": "126144000s",
            "amount": 40000000000000,
            "amount_multiplier": "0.500000000000000000"
          }
        }
      ]
    }
  }
}
```

2. Linear minting of 100 million of token during period of 10 years, next no emission

```json
{
  "params": {
    "mint_denom": "uc4e",
    "minter_config": {
      "start_time": "2022-07-05T00:00:00Z",
      "minters": [
        {
          "sequenceId": 1,
          "end_time": "2023-07-05T00:00:00Z",
          "type": "LINEAR_MINTING",
          "linear_minting": {
            "amount": 100000000000000
          },
          "exponential_step_minting": null
        },
        {
          "sequenceId": 2,
          "end_time": null,
          "type": "NO_MINTING",
          "linear_minting": null,
          "exponential_step_minting": null
        }
      ]
    }
  }
}
```

## State

Chain4Energy minter module state contains information used by current minter.
Module state contains following data:

| Key                | Type                         | Description                   |
|--------------------|------------------------------|-------------------------------|
| minter_state       | MinterState                  | current minter state          |
| state_history      | List of MinterState          | previous minters final states |

### MinterState

| Key                            | Type     | Description                                                        |
|--------------------------------|----------|--------------------------------------------------------------------|
| sequence_id                    | uint32   | current minter sequenceId                                          |
| amount_minted                  | sdk.Int  | amount minted by current minter                                    |
| remainder_to_mint              | sdk.Dec  | amount that should have been minted in previous block but was not  |
| last_mint_block_time           | sdk.Time | Time of last mint                                                  |
| remainder_from_previous_period | sdk.Dec  | amount that should have been minted in previous minter but was not | // TODO: rename to remainder_from_previous_minter

### Example state

```json
{
  "minter_state": {
    "sequence_id": 1,
    "amount_minted": "13766330043442",
    "remainder_to_mint": "0.415017757483510908",
    "last_mint_block_time": "2022-11-07T14:49:34.606250Z",
    "remainder_from_previous_period": "0.000000000000000000"
  },
  "state_history": []
}
```

## Events

Chain4Energy minter module emits the following events:

### BeginBlockers

#### Mint

| Type         | Attribute Key | Attribute Value               |
|--------------|---------------|-------------------------------|
| bonded_ratio | Dec           |                               |
| inflation    | sdk.Dec       | Minting block inflation level |
| amount       | sdk.Int       | Amount minted in block        |

TODO remove bonded_ratio 
TODO add minter period id

## Queries

### Params query

Queries the module params.

See example response:

```json
{
  "params": {
    "mint_denom": "uc4e",
    "minter_config": {
      "start_time": "2022-07-05T00:00:00Z",
      "minters": [
        {
          "sequence_id": 1,
          "end_time": null,
          "type": "EXPONENTIAL_STEP_MINTING",
          "linear_minting": null,
          "exponential_step_minting": {
            "step_duration": "31536000s",
            "amount": "40000000000000",
            "amount_multiplier": "0.500000000000000000"
          }
        }
      ]
    }
  }
}
```
### State query

Queries the module state.

See example response:

```json
{
  "minter_state": {
    "sequence_id": 1,
    "amount_minted": "13766330043442",
    "remainder_to_mint": "0.415017757483510908",
    "last_mint_block_time": "2022-11-07T14:49:34.606250Z",
    "remainder_from_previous_period": "0.000000000000000000"
  },
  "state_history": []
}
```

### Inflation query

Queries current inflation.

See example response:

```json
{
  "inflation": "0.102489480201216908"
}
```

## Genesis validations

TODO
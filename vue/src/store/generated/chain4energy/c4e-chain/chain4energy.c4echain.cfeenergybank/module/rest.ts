/* eslint-disable */
/* tslint:disable */
/*
 * ---------------------------------------------------------------
 * ## THIS FILE WAS GENERATED VIA SWAGGER-TYPESCRIPT-API        ##
 * ##                                                           ##
 * ## AUTHOR: acacode                                           ##
 * ## SOURCE: https://github.com/acacode/swagger-typescript-api ##
 * ---------------------------------------------------------------
 */

export interface CfeenergybankEnergyToken {
  /** @format uint64 */
  id?: string;
  name?: string;

  /** @format uint64 */
  amount?: string;
  userAddress?: string;

  /** @format uint64 */
  createdAt?: string;
}

export type CfeenergybankMsgCreateTokenParamsResponse = object;

export type CfeenergybankMsgMintTokenResponse = object;

export type CfeenergybankMsgTransferTokensOptimallyResponse = object;

export type CfeenergybankMsgTransferTokensResponse = object;

/**
 * Params defines the parameters for the module.
 */
export type CfeenergybankParams = object;

export interface CfeenergybankQueryAllEnergyTokenResponse {
  EnergyToken?: CfeenergybankEnergyToken[];

  /**
   * PageResponse is to be embedded in gRPC response messages where the
   * corresponding request message has used PageRequest.
   *
   *  message SomeResponse {
   *          repeated Bar results = 1;
   *          PageResponse page = 2;
   *  }
   */
  pagination?: V1Beta1PageResponse;
}

export interface CfeenergybankQueryAllTokenParamsResponse {
  tokenParams?: CfeenergybankTokenParams[];

  /**
   * PageResponse is to be embedded in gRPC response messages where the
   * corresponding request message has used PageRequest.
   *
   *  message SomeResponse {
   *          repeated Bar results = 1;
   *          PageResponse page = 2;
   *  }
   */
  pagination?: V1Beta1PageResponse;
}

export interface CfeenergybankQueryAllTokensHistoryResponse {
  TokensHistory?: CfeenergybankTokensHistory[];

  /**
   * PageResponse is to be embedded in gRPC response messages where the
   * corresponding request message has used PageRequest.
   *
   *  message SomeResponse {
   *          repeated Bar results = 1;
   *          PageResponse page = 2;
   *  }
   */
  pagination?: V1Beta1PageResponse;
}

export interface CfeenergybankQueryCurrentBalanceResponse {
  /** @format uint64 */
  balance?: string;
}

export interface CfeenergybankQueryEnergyTokenUserAddressResponse {
  EnergyToken?: CfeenergybankEnergyToken[];

  /**
   * PageResponse is to be embedded in gRPC response messages where the
   * corresponding request message has used PageRequest.
   *
   *  message SomeResponse {
   *          repeated Bar results = 1;
   *          PageResponse page = 2;
   *  }
   */
  pagination?: V1Beta1PageResponse;
}

export interface CfeenergybankQueryGetEnergyTokenResponse {
  EnergyToken?: CfeenergybankEnergyToken;
}

export interface CfeenergybankQueryGetTokenParamsResponse {
  tokenParams?: CfeenergybankTokenParams;
}

export interface CfeenergybankQueryGetTokensHistoryResponse {
  TokensHistory?: CfeenergybankTokensHistory;
}

/**
 * QueryParamsResponse is response type for the Query/Params RPC method.
 */
export interface CfeenergybankQueryParamsResponse {
  /** params holds all the parameters of this module. */
  params?: CfeenergybankParams;
}

export interface CfeenergybankQueryTokensHistoryUserAddressResponse {
  TokensHistory?: CfeenergybankTokensHistory[];
}

export interface CfeenergybankTokenParams {
  index?: string;
  name?: string;
  tradingCompany?: string;

  /** @format uint64 */
  burningTime?: string;
  burningType?: string;
  mintAccount?: string;

  /** @format uint64 */
  exchangeRate?: string;

  /** @format uint64 */
  commissionRate?: string;
}

export interface CfeenergybankTokensHistory {
  /** @format uint64 */
  id?: string;
  userAddress?: string;

  /** @format uint64 */
  createdAt?: string;
  issuerAddress?: string;
  targetAddress?: string;

  /** @format uint64 */
  amount?: string;
  tokenName?: string;
  operationType?: string;
}

export interface ProtobufAny {
  "@type"?: string;
}

export interface RpcStatus {
  /** @format int32 */
  code?: number;
  message?: string;
  details?: ProtobufAny[];
}

/**
* message SomeRequest {
         Foo some_parameter = 1;
         PageRequest pagination = 2;
 }
*/
export interface V1Beta1PageRequest {
  /**
   * key is a value returned in PageResponse.next_key to begin
   * querying the next page most efficiently. Only one of offset or key
   * should be set.
   * @format byte
   */
  key?: string;

  /**
   * offset is a numeric offset that can be used when key is unavailable.
   * It is less efficient than using key. Only one of offset or key should
   * be set.
   * @format uint64
   */
  offset?: string;

  /**
   * limit is the total number of results to be returned in the result page.
   * If left empty it will default to a value to be set by each app.
   * @format uint64
   */
  limit?: string;

  /**
   * count_total is set to true  to indicate that the result set should include
   * a count of the total number of items available for pagination in UIs.
   * count_total is only respected when offset is used. It is ignored when key
   * is set.
   */
  count_total?: boolean;
}

/**
* PageResponse is to be embedded in gRPC response messages where the
corresponding request message has used PageRequest.

 message SomeResponse {
         repeated Bar results = 1;
         PageResponse page = 2;
 }
*/
export interface V1Beta1PageResponse {
  /** @format byte */
  next_key?: string;

  /** @format uint64 */
  total?: string;
}

export type QueryParamsType = Record<string | number, any>;
export type ResponseFormat = keyof Omit<Body, "body" | "bodyUsed">;

export interface FullRequestParams extends Omit<RequestInit, "body"> {
  /** set parameter to `true` for call `securityWorker` for this request */
  secure?: boolean;
  /** request path */
  path: string;
  /** content type of request body */
  type?: ContentType;
  /** query params */
  query?: QueryParamsType;
  /** format of response (i.e. response.json() -> format: "json") */
  format?: keyof Omit<Body, "body" | "bodyUsed">;
  /** request body */
  body?: unknown;
  /** base url */
  baseUrl?: string;
  /** request cancellation token */
  cancelToken?: CancelToken;
}

export type RequestParams = Omit<FullRequestParams, "body" | "method" | "query" | "path">;

export interface ApiConfig<SecurityDataType = unknown> {
  baseUrl?: string;
  baseApiParams?: Omit<RequestParams, "baseUrl" | "cancelToken" | "signal">;
  securityWorker?: (securityData: SecurityDataType) => RequestParams | void;
}

export interface HttpResponse<D extends unknown, E extends unknown = unknown> extends Response {
  data: D;
  error: E;
}

type CancelToken = Symbol | string | number;

export enum ContentType {
  Json = "application/json",
  FormData = "multipart/form-data",
  UrlEncoded = "application/x-www-form-urlencoded",
}

export class HttpClient<SecurityDataType = unknown> {
  public baseUrl: string = "";
  private securityData: SecurityDataType = null as any;
  private securityWorker: null | ApiConfig<SecurityDataType>["securityWorker"] = null;
  private abortControllers = new Map<CancelToken, AbortController>();

  private baseApiParams: RequestParams = {
    credentials: "same-origin",
    headers: {},
    redirect: "follow",
    referrerPolicy: "no-referrer",
  };

  constructor(apiConfig: ApiConfig<SecurityDataType> = {}) {
    Object.assign(this, apiConfig);
  }

  public setSecurityData = (data: SecurityDataType) => {
    this.securityData = data;
  };

  private addQueryParam(query: QueryParamsType, key: string) {
    const value = query[key];

    return (
      encodeURIComponent(key) +
      "=" +
      encodeURIComponent(Array.isArray(value) ? value.join(",") : typeof value === "number" ? value : `${value}`)
    );
  }

  protected toQueryString(rawQuery?: QueryParamsType): string {
    const query = rawQuery || {};
    const keys = Object.keys(query).filter((key) => "undefined" !== typeof query[key]);
    return keys
      .map((key) =>
        typeof query[key] === "object" && !Array.isArray(query[key])
          ? this.toQueryString(query[key] as QueryParamsType)
          : this.addQueryParam(query, key),
      )
      .join("&");
  }

  protected addQueryParams(rawQuery?: QueryParamsType): string {
    const queryString = this.toQueryString(rawQuery);
    return queryString ? `?${queryString}` : "";
  }

  private contentFormatters: Record<ContentType, (input: any) => any> = {
    [ContentType.Json]: (input: any) =>
      input !== null && (typeof input === "object" || typeof input === "string") ? JSON.stringify(input) : input,
    [ContentType.FormData]: (input: any) =>
      Object.keys(input || {}).reduce((data, key) => {
        data.append(key, input[key]);
        return data;
      }, new FormData()),
    [ContentType.UrlEncoded]: (input: any) => this.toQueryString(input),
  };

  private mergeRequestParams(params1: RequestParams, params2?: RequestParams): RequestParams {
    return {
      ...this.baseApiParams,
      ...params1,
      ...(params2 || {}),
      headers: {
        ...(this.baseApiParams.headers || {}),
        ...(params1.headers || {}),
        ...((params2 && params2.headers) || {}),
      },
    };
  }

  private createAbortSignal = (cancelToken: CancelToken): AbortSignal | undefined => {
    if (this.abortControllers.has(cancelToken)) {
      const abortController = this.abortControllers.get(cancelToken);
      if (abortController) {
        return abortController.signal;
      }
      return void 0;
    }

    const abortController = new AbortController();
    this.abortControllers.set(cancelToken, abortController);
    return abortController.signal;
  };

  public abortRequest = (cancelToken: CancelToken) => {
    const abortController = this.abortControllers.get(cancelToken);

    if (abortController) {
      abortController.abort();
      this.abortControllers.delete(cancelToken);
    }
  };

  public request = <T = any, E = any>({
    body,
    secure,
    path,
    type,
    query,
    format = "json",
    baseUrl,
    cancelToken,
    ...params
  }: FullRequestParams): Promise<HttpResponse<T, E>> => {
    const secureParams = (secure && this.securityWorker && this.securityWorker(this.securityData)) || {};
    const requestParams = this.mergeRequestParams(params, secureParams);
    const queryString = query && this.toQueryString(query);
    const payloadFormatter = this.contentFormatters[type || ContentType.Json];

    return fetch(`${baseUrl || this.baseUrl || ""}${path}${queryString ? `?${queryString}` : ""}`, {
      ...requestParams,
      headers: {
        ...(type && type !== ContentType.FormData ? { "Content-Type": type } : {}),
        ...(requestParams.headers || {}),
      },
      signal: cancelToken ? this.createAbortSignal(cancelToken) : void 0,
      body: typeof body === "undefined" || body === null ? null : payloadFormatter(body),
    }).then(async (response) => {
      const r = response as HttpResponse<T, E>;
      r.data = (null as unknown) as T;
      r.error = (null as unknown) as E;

      const data = await response[format]()
        .then((data) => {
          if (r.ok) {
            r.data = data;
          } else {
            r.error = data;
          }
          return r;
        })
        .catch((e) => {
          r.error = e;
          return r;
        });

      if (cancelToken) {
        this.abortControllers.delete(cancelToken);
      }

      if (!response.ok) throw data;
      return data;
    });
  };
}

/**
 * @title cfeenergybank/energy_token.proto
 * @version version not set
 */
export class Api<SecurityDataType extends unknown> extends HttpClient<SecurityDataType> {
  /**
   * No description
   *
   * @tags Query
   * @name QueryCurrentBalance
   * @summary Queries a list of CurrentBalance items.
   * @request GET:/chain4energy/c4e-chain/cfeenergybank/current_balance/{userAddress}/{tokenName}
   */
  queryCurrentBalance = (userAddress: string, tokenName: string, params: RequestParams = {}) =>
    this.request<CfeenergybankQueryCurrentBalanceResponse, RpcStatus>({
      path: `/chain4energy/c4e-chain/cfeenergybank/current_balance/${userAddress}/${tokenName}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryEnergyTokenAll
   * @summary Queries a list of EnergyToken items.
   * @request GET:/chain4energy/c4e-chain/cfeenergybank/energy_token
   */
  queryEnergyTokenAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<CfeenergybankQueryAllEnergyTokenResponse, RpcStatus>({
      path: `/chain4energy/c4e-chain/cfeenergybank/energy_token`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryEnergyToken
   * @summary Queries a EnergyToken by id.
   * @request GET:/chain4energy/c4e-chain/cfeenergybank/energy_token/{id}
   */
  queryEnergyToken = (id: string, params: RequestParams = {}) =>
    this.request<CfeenergybankQueryGetEnergyTokenResponse, RpcStatus>({
      path: `/chain4energy/c4e-chain/cfeenergybank/energy_token/${id}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryEnergyTokenUserAddress
   * @summary Queries a list of EnergyTokenUserAddress items.
   * @request GET:/chain4energy/c4e-chain/cfeenergybank/energy_token_user_address/{userAddress}
   */
  queryEnergyTokenUserAddress = (
    userAddress: string,
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<CfeenergybankQueryEnergyTokenUserAddressResponse, RpcStatus>({
      path: `/chain4energy/c4e-chain/cfeenergybank/energy_token_user_address/${userAddress}`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryParams
   * @summary Parameters queries the parameters of the module.
   * @request GET:/chain4energy/c4e-chain/cfeenergybank/params
   */
  queryParams = (params: RequestParams = {}) =>
    this.request<CfeenergybankQueryParamsResponse, RpcStatus>({
      path: `/chain4energy/c4e-chain/cfeenergybank/params`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryTokenParamsAll
   * @summary Queries a list of TokenParams items.
   * @request GET:/chain4energy/c4e-chain/cfeenergybank/token_params
   */
  queryTokenParamsAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<CfeenergybankQueryAllTokenParamsResponse, RpcStatus>({
      path: `/chain4energy/c4e-chain/cfeenergybank/token_params`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryTokenParams
   * @summary Queries a TokenParams by index.
   * @request GET:/chain4energy/c4e-chain/cfeenergybank/token_params/{index}
   */
  queryTokenParams = (index: string, params: RequestParams = {}) =>
    this.request<CfeenergybankQueryGetTokenParamsResponse, RpcStatus>({
      path: `/chain4energy/c4e-chain/cfeenergybank/token_params/${index}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryTokensHistoryAll
   * @summary Queries a list of TokensHistory items.
   * @request GET:/chain4energy/c4e-chain/cfeenergybank/tokens_history
   */
  queryTokensHistoryAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<CfeenergybankQueryAllTokensHistoryResponse, RpcStatus>({
      path: `/chain4energy/c4e-chain/cfeenergybank/tokens_history`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryTokensHistory
   * @summary Queries a TokensHistory by id.
   * @request GET:/chain4energy/c4e-chain/cfeenergybank/tokens_history/{id}
   */
  queryTokensHistory = (id: string, params: RequestParams = {}) =>
    this.request<CfeenergybankQueryGetTokensHistoryResponse, RpcStatus>({
      path: `/chain4energy/c4e-chain/cfeenergybank/tokens_history/${id}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryTokensHistoryUserAddress
   * @summary Queries a list of TokensHistoryUserAddress items.
   * @request GET:/chain4energy/c4e-chain/cfeenergybank/tokens_history_user_address/{userBlockchainAddress}
   */
  queryTokensHistoryUserAddress = (userBlockchainAddress: string, params: RequestParams = {}) =>
    this.request<CfeenergybankQueryTokensHistoryUserAddressResponse, RpcStatus>({
      path: `/chain4energy/c4e-chain/cfeenergybank/tokens_history_user_address/${userBlockchainAddress}`,
      method: "GET",
      format: "json",
      ...params,
    });
}

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

export interface CfeminterMinter {
  /** @format date-time */
  start?: string;
  periods?: CfeminterMintingPeriod[];
}

export interface CfeminterMinterState {
  /** @format int32 */
  position?: number;
  amount_minted?: string;
  remainder_to_mint?: string;

  /** @format date-time */
  last_mint_block_time?: string;
  remainder_from_previous_period?: string;
}

export interface CfeminterMintingPeriod {
  /** @format int32 */
  position?: number;

  /** @format date-time */
  period_end?: string;

  /**
   * types:
   *   NO_MINTING;
   *   TIME_LINEAR_MINTER;
   *   PERIODIC_REDUCTION_MINTER;
   */
  type?: string;
  time_linear_minter?: CfeminterTimeLinearMinter;
  periodic_reduction_minter?: CfeminterPeriodicReductionMinter;
}

/**
 * Params defines the parameters for the module.
 */
export interface CfeminterParams {
  mint_denom?: string;
  minter?: CfeminterMinter;
}

export interface CfeminterPeriodicReductionMinter {
  /**
   * mint_period in seconds
   * @format int32
   */
  mint_period?: number;
  mint_amount?: string;

  /** @format int32 */
  reduction_period_length?: number;
  reduction_factor?: string;
}

export interface CfeminterQueryInflationResponse {
  inflation?: string;
}

/**
 * QueryParamsResponse is response type for the Query/Params RPC method.
 */
export interface CfeminterQueryParamsResponse {
  /** Params defines the parameters for the module. */
  params?: CfeminterParams;
}

export interface CfeminterQueryStateResponse {
  minter_state?: CfeminterMinterState;
  state_history?: CfeminterMinterState[];
}

export interface CfeminterTimeLinearMinter {
  amount?: string;
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

import axios, { AxiosInstance, AxiosRequestConfig, AxiosResponse, ResponseType } from "axios";

export type QueryParamsType = Record<string | number, any>;

export interface FullRequestParams extends Omit<AxiosRequestConfig, "data" | "params" | "url" | "responseType"> {
  /** set parameter to `true` for call `securityWorker` for this request */
  secure?: boolean;
  /** request path */
  path: string;
  /** content type of request body */
  type?: ContentType;
  /** query params */
  query?: QueryParamsType;
  /** format of response (i.e. response.json() -> format: "json") */
  format?: ResponseType;
  /** request body */
  body?: unknown;
}

export type RequestParams = Omit<FullRequestParams, "body" | "method" | "query" | "path">;

export interface ApiConfig<SecurityDataType = unknown> extends Omit<AxiosRequestConfig, "data" | "cancelToken"> {
  securityWorker?: (
    securityData: SecurityDataType | null,
  ) => Promise<AxiosRequestConfig | void> | AxiosRequestConfig | void;
  secure?: boolean;
  format?: ResponseType;
}

export enum ContentType {
  Json = "application/json",
  FormData = "multipart/form-data",
  UrlEncoded = "application/x-www-form-urlencoded",
}

export class HttpClient<SecurityDataType = unknown> {
  public instance: AxiosInstance;
  private securityData: SecurityDataType | null = null;
  private securityWorker?: ApiConfig<SecurityDataType>["securityWorker"];
  private secure?: boolean;
  private format?: ResponseType;

  constructor({ securityWorker, secure, format, ...axiosConfig }: ApiConfig<SecurityDataType> = {}) {
    this.instance = axios.create({ ...axiosConfig, baseURL: axiosConfig.baseURL || "" });
    this.secure = secure;
    this.format = format;
    this.securityWorker = securityWorker;
  }

  public setSecurityData = (data: SecurityDataType | null) => {
    this.securityData = data;
  };

  private mergeRequestParams(params1: AxiosRequestConfig, params2?: AxiosRequestConfig): AxiosRequestConfig {
    return {
      ...this.instance.defaults,
      ...params1,
      ...(params2 || {}),
      headers: {
        ...(this.instance.defaults.headers || {}),
        ...(params1.headers || {}),
        ...((params2 && params2.headers) || {}),
      },
    };
  }

  private createFormData(input: Record<string, unknown>): FormData {
    return Object.keys(input || {}).reduce((formData, key) => {
      const property = input[key];
      formData.append(
        key,
        property instanceof Blob
          ? property
          : typeof property === "object" && property !== null
          ? JSON.stringify(property)
          : `${property}`,
      );
      return formData;
    }, new FormData());
  }

  public request = async <T = any, _E = any>({
    secure,
    path,
    type,
    query,
    format,
    body,
    ...params
  }: FullRequestParams): Promise<AxiosResponse<T>> => {
    const secureParams =
      ((typeof secure === "boolean" ? secure : this.secure) &&
        this.securityWorker &&
        (await this.securityWorker(this.securityData))) ||
      {};
    const requestParams = this.mergeRequestParams(params, secureParams);
    const responseFormat = (format && this.format) || void 0;

    if (type === ContentType.FormData && body && body !== null && typeof body === "object") {
      requestParams.headers.common = { Accept: "*/*" };
      requestParams.headers.post = {};
      requestParams.headers.put = {};

      body = this.createFormData(body as Record<string, unknown>);
    }

    return this.instance.request({
      ...requestParams,
      headers: {
        ...(type && type !== ContentType.FormData ? { "Content-Type": type } : {}),
        ...(requestParams.headers || {}),
      },
      params: query,
      responseType: responseFormat,
      data: body,
      url: path,
    });
  };
}

/**
 * @title c4e-chain/cfeminter/event.proto
 * @version version not set
 */
export class Api<SecurityDataType extends unknown> extends HttpClient<SecurityDataType> {
  /**
   * No description
   *
   * @tags Query
   * @name QueryInflation
   * @summary Queries a list of Inflation items.
   * @request GET:/c4e/minter/inflation
   */
  queryInflation = (params: RequestParams = {}) =>
    this.request<CfeminterQueryInflationResponse, RpcStatus>({
      path: `/c4e/minter/inflation`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryParams
   * @summary Parameters queries the parameters of the module.
   * @request GET:/c4e/minter/params
   */
  queryParams = (params: RequestParams = {}) =>
    this.request<CfeminterQueryParamsResponse, RpcStatus>({
      path: `/c4e/minter/params`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryState
   * @summary Queries a list of State items.
   * @request GET:/c4e/minter/state
   */
  queryState = (params: RequestParams = {}) =>
    this.request<CfeminterQueryStateResponse, RpcStatus>({
      path: `/c4e/minter/state`,
      method: "GET",
      format: "json",
      ...params,
    });
}

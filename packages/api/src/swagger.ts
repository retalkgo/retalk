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

export interface CommonResp {
	data?: any;
	msg?: string;
	success?: boolean;
}

export interface EntityAuthor {
	created_at?: string;
	email?: string;
	id?: number;
	is_admin?: boolean;
	link?: string;
	name?: string;
}

export interface EntityCookedComment {
	author?: EntityAuthor;
	body?: string;
	created_at?: string;
	id?: number;
	path?: string;
}

export interface HandlerRespInit {
	isInit?: boolean;
}

export type QueryParamsType = Record<string | number, any>;
export type ResponseFormat = keyof Omit<Body, "body" | "bodyUsed">;

export interface FullRequestParams extends Omit<RequestInit, "body"> {
	/**
	 * Set parameter to `true` for call `securityWorker` for this request
	 */
	secure?: boolean;
	/**
	 * Request path
	 */
	path: string;
	/**
	 * Content type of request body
	 */
	type?: ContentType;
	/**
	 * Query params
	 */
	query?: QueryParamsType;
	/**
	 * Format of response (i.e. response.json() -> format: "json")
	 */
	format?: ResponseFormat;
	/**
	 * Request body
	 */
	body?: unknown;
	/**
	 * Base url
	 */
	baseUrl?: string;
	/**
	 * Request cancellation token
	 */
	cancelToken?: CancelToken;
}

export type RequestParams = Omit<
	FullRequestParams,
	"body" | "method" | "query" | "path"
>;

export interface ApiConfig<SecurityDataType = unknown> {
	baseUrl?: string;
	baseApiParams?: Omit<RequestParams, "baseUrl" | "cancelToken" | "signal">;
	securityWorker?: (
		securityData: SecurityDataType | null,
	) => Promise<RequestParams | void> | RequestParams | void;
	customFetch?: typeof fetch;
}

export interface HttpResponse<D extends unknown, E extends unknown = unknown>
	extends Response {
	data: D;
	error: E;
}

type CancelToken = Symbol | string | number;

export enum ContentType {
	Json = "application/json",
	FormData = "multipart/form-data",
	UrlEncoded = "application/x-www-form-urlencoded",
	Text = "text/plain",
}

export class HttpClient<SecurityDataType = unknown> {
	public baseUrl: string = "/";
	private securityData: SecurityDataType | null = null;
	private securityWorker?: ApiConfig<SecurityDataType>["securityWorker"];
	private abortControllers = new Map<CancelToken, AbortController>();
	private customFetch = (...fetchParams: Parameters<typeof fetch>) =>
		fetch(...fetchParams);

	private baseApiParams: RequestParams = {
		credentials: "same-origin",
		headers: {},
		redirect: "follow",
		referrerPolicy: "no-referrer",
	};

	constructor(apiConfig: ApiConfig<SecurityDataType> = {}) {
		Object.assign(this, apiConfig);
	}

	public setSecurityData = (data: SecurityDataType | null) => {
		this.securityData = data;
	};

	protected encodeQueryParam(key: string, value: any) {
		const encodedKey = encodeURIComponent(key);
		return `${encodedKey}=${encodeURIComponent(
			typeof value === "number" ? value : `${value}`,
		)}`;
	}

	protected addQueryParam(query: QueryParamsType, key: string) {
		return this.encodeQueryParam(key, query[key]);
	}

	protected addArrayQueryParam(query: QueryParamsType, key: string) {
		const value = query[key];
		return value.map((v: any) => this.encodeQueryParam(key, v)).join("&");
	}

	protected toQueryString(rawQuery?: QueryParamsType): string {
		const query = rawQuery || {};
		const keys = Object.keys(query).filter(
			(key) => "undefined" !== typeof query[key],
		);
		return keys
			.map((key) =>
				Array.isArray(query[key])
					? this.addArrayQueryParam(query, key)
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
			input !== null && (typeof input === "object" || typeof input === "string")
				? JSON.stringify(input)
				: input,
		[ContentType.Text]: (input: any) =>
			input !== null && typeof input !== "string"
				? JSON.stringify(input)
				: input,
		[ContentType.FormData]: (input: any) =>
			Object.keys(input || {}).reduce((formData, key) => {
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
			}, new FormData()),
		[ContentType.UrlEncoded]: (input: any) => this.toQueryString(input),
	};

	protected mergeRequestParams(
		params1: RequestParams,
		params2?: RequestParams,
	): RequestParams {
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

	protected createAbortSignal = (
		cancelToken: CancelToken,
	): AbortSignal | undefined => {
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

	public request = async <T = any, E = any>({
		body,
		secure,
		path,
		type,
		query,
		format,
		baseUrl,
		cancelToken,
		...params
	}: FullRequestParams): Promise<HttpResponse<T, E>> => {
		const secureParams =
			((typeof secure === "boolean" ? secure : this.baseApiParams.secure) &&
				this.securityWorker &&
				(await this.securityWorker(this.securityData))) ||
			{};
		const requestParams = this.mergeRequestParams(params, secureParams);
		const queryString = query && this.toQueryString(query);
		const payloadFormatter = this.contentFormatters[type || ContentType.Json];
		const responseFormat = format || requestParams.format;

		return this.customFetch(
			`${baseUrl || this.baseUrl || ""}${path}${
				queryString ? `?${queryString}` : ""
			}`,
			{
				...requestParams,
				headers: {
					...(requestParams.headers || {}),
					...(type && type !== ContentType.FormData
						? { "Content-Type": type }
						: {}),
				},
				signal:
					(cancelToken
						? this.createAbortSignal(cancelToken)
						: requestParams.signal) || null,
				body:
					typeof body === "undefined" || body === null
						? null
						: payloadFormatter(body),
			},
		).then(async (response) => {
			const r = response as HttpResponse<T, E>;
			r.data = null as unknown as T;
			r.error = null as unknown as E;

			const data = !responseFormat
				? r
				: await response[responseFormat]()
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

			if (!response.ok) {
				throw data;
			}
			return data;
		});
	};
}

/**
 * @license GPL-3.0
 * @version 1.0
 * @title Retalk API
 * @baseUrl /
 * @contact API 支持 <retalk@redish101.top>
 *
 * Retalk 后端 API 文档
 */
export class Api<
	SecurityDataType extends unknown,
> extends HttpClient<SecurityDataType> {
	/**
	 * 输出欢迎信息以验证安装
	 *
	 * @tags 首页
	 * @name GetRoot
	 * @summary 首页
	 * @request GET:/
	 */
	getRoot = (params: RequestParams = {}) =>
		this.request<
			CommonResp & {
				msg?: string;
			},
			any
		>({
			path: `/`,
			method: "GET",
			...params,
		});

	api = {
		/**
		 * 新增评论
		 *
		 * @tags 评论
		 * @name CommentAddCreate
		 * @summary 新增评论
		 * @request POST:/api/comment/add
		 */
		commentAddCreate: (
			data: {
				/**
				 * 评论路径
				 */
				path: string;
				/**
				 * 发送者昵称
				 */
				name: string;
				/**
				 * 发送者邮箱
				 */
				email: string;
				/**
				 * 发送者网站
				 */
				link: string;
				/**
				 * 正文
				 */
				body: string;
			},
			params: RequestParams = {},
		) =>
			this.request<CommonResp, CommonResp>({
				path: `/api/comment/add`,
				method: "POST",
				body: data,
				type: ContentType.UrlEncoded,
				...params,
			}),

		/**
		 * 根据ID删除评论
		 *
		 * @tags 评论
		 * @name CommentDeleteDelete
		 * @summary 根据ID删除评论
		 * @request DELETE:/api/comment/delete
		 * @secure
		 */
		commentDeleteDelete: (
			query: {
				/**
				 * 评论ID
				 */
				id: string;
			},
			params: RequestParams = {},
		) =>
			this.request<CommonResp, CommonResp>({
				path: `/api/comment/delete`,
				method: "DELETE",
				query: query,
				secure: true,
				...params,
			}),

		/**
		 * 获取所有评论
		 *
		 * @tags 评论
		 * @name CommentGetAllList
		 * @summary 获取所有评论
		 * @request GET:/api/comment/getAll
		 */
		commentGetAllList: (params: RequestParams = {}) =>
			this.request<
				CommonResp & {
					data?: EntityCookedComment[];
				},
				CommonResp
			>({
				path: `/api/comment/getAll`,
				method: "GET",
				...params,
			}),

		/**
		 * 根据路径获取评论
		 *
		 * @tags 评论
		 * @name CommentGetByPathList
		 * @summary 根据路径获取评论
		 * @request GET:/api/comment/getByPath
		 */
		commentGetByPathList: (
			query: {
				/**
				 * 路径
				 */
				path: string;
			},
			params: RequestParams = {},
		) =>
			this.request<
				CommonResp & {
					data?: EntityCookedComment[];
				},
				CommonResp
			>({
				path: `/api/comment/getByPath`,
				method: "GET",
				query: query,
				...params,
			}),

		/**
		 * 初始化服务端, 创建ApiKey
		 *
		 * @tags 服务端
		 * @name InitCreate
		 * @summary 初始化服务端
		 * @request POST:/api/init
		 */
		initCreate: (
			data: {
				/**
				 * ApiKey
				 */
				apikey: string;
				/**
				 * 管理员昵称
				 */
				name: string;
				/**
				 * 管理员邮箱
				 */
				email: string;
				/**
				 * 管理员网站
				 */
				link: string;
			},
			params: RequestParams = {},
		) =>
			this.request<
				CommonResp & {
					data?: HandlerRespInit;
				},
				| (CommonResp & {
						data?: HandlerRespInit;
				  })
				| CommonResp
			>({
				path: `/api/init`,
				method: "POST",
				body: data,
				type: ContentType.UrlEncoded,
				...params,
			}),
	};
}

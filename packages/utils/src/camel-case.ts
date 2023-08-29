type CamelCase<T extends string> =
	T extends `${infer P1}_${infer P2}${infer P3}`
		? `${Lowercase<P1>}${Uppercase<P2>}${CamelCase<P3>}`
		: Lowercase<T>;

type ToCamelCase<T extends Record<string, any>> = {
	[K in keyof T as CamelCase<K & string>]: T[K];
};

type Simplify<T> = { [K in keyof T]: Simplify<T[K]> } & {};

export type DeepToCamelCase<T> = Simplify<
	T extends Record<string, any>
		? ToCamelCase<{ [K in keyof T]: DeepToCamelCase<T[K]> }>
		: T
>;
export function deepToCamelCase<T>(obj: T): DeepToCamelCase<T> {
	if (Array.isArray(obj)) {
		return obj.map((item) => deepToCamelCase(item)) as any;
	}
	if (obj !== null && typeof obj === "object") {
		return Object.fromEntries(
			Object.entries(obj).map(([key, value]) => [
				key.replace(/_[a-z]/g, (g) => g[1].toUpperCase()),
				deepToCamelCase(value),
			]),
		) as any;
	}

	return obj as any;
}

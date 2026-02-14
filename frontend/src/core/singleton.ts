export type ClassLike<T> = object & { prototype: T };

export default abstract class Singleton {
  private static _instances = new WeakMap<object, unknown>();

  protected constructor() {}

  static getInstance<T>(this: ClassLike<T>): T {
    const key = this as object;

    let instance = Singleton._instances.get(key) as T | undefined;
    if (!instance) {
      const Self = this as unknown as new () => T;
      instance = new Self();

      Singleton._instances.set(key, instance);
    }

    return instance;
  }
}

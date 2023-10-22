import { mul } from "./mul";
import { mul2 } from "./mul2";

export function add(x: number, y: number) {
  return x + y;
}

export function mul_and_add(mulX: number, mulY: number, addX: number) {
  return mul(mulX, mulY) + addX;
}

export function mul_and_add2(mulX: number, mulY: number, addX: number) {
  return mul2(mulX, mulY) + addX;
}

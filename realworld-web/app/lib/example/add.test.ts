import { add, mul_and_add, mul_and_add2 } from "./add";
import * as mul from "./mul";

jest.mock("./mul");

const mulMock = jest.mocked(mul);

describe("add", () => {
  it("given 1 and 2 return 3", () => {
    expect(add(1, 2)).toBe(3);
  });
});

describe("mul_and_add", () => {
  describe("given 1, 2 and 3", () => {
    it("return 5", () => {
      mulMock.mul.mockReturnValueOnce(2);
      const got = mul_and_add(1, 2, 3);

      expect(got).toBe(5);
      expect(mulMock.mul).toHaveBeenCalledWith(1, 2);
    });

    it("return 6", () => {
      const got = mul_and_add(1, 2, 3);
      expect(got).toBe(6);
      expect(mulMock.mul).toHaveBeenCalledWith(1, 2);
    });
  });
});

describe("mul_and_add2", () => {
  describe("given 1, 2 and 3", () => {});
  it("return 5", () => {
    const got = mul_and_add2(1, 2, 3);
    expect(got).toBe(5);
  });
});

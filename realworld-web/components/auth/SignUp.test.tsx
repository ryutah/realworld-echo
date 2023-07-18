import { render } from "@testing-library/react";
import SignUp from "@/components/auth/SignUp";

describe("SignUp", () => {
  it("should render the SignUp component", () => {
    const { getByTestId } = render(<SignUp testIds={{ signup: "signup" }} />);
    expect(getByTestId("signup")).toBeTruthy();
  });
});

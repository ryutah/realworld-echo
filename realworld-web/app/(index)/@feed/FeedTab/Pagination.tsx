import { Pagination as MuiPagination } from "@mui/material";
import { useRouter, useSearchParams } from "next/navigation";

type Props = {
  count: number;
};

export default function Pagination({ count }: Props) {
  const params = useSearchParams();
  const router = useRouter();

  const page = params.get("page") ? Number(params.get("page")) : 1;

  return (
    <MuiPagination
      size="large"
      count={count}
      page={page}
      sx={{ textAlign: "center", alignItems: "center" }}
      onChange={(_, number) => router.push(`/?page=${number}`)}
    />
  );
}

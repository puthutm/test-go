"use client";

import { useEffect } from "react";
import { Button, Card, CardBody } from "reactstrap";

export default function Error({
  error,
  reset,
}: {
  error: Error & { digest?: string };
  reset: () => void;
}) {
  useEffect(() => {
    console.error(error);
  }, [error]);

  return (
    <Card className="mt-3">
      <CardBody>
        <h2>
          {process.env.NODE_ENV === "production"
            ? "Something went wrong"
            : error.message}
        </h2>
        <Button
          color="primary"
          onClick={
            // Attempt to recover by trying to re-render the segment
            () => reset()
          }
        >
          Try again
        </Button>
      </CardBody>
    </Card>
  );
}

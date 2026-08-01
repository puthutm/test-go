import { Card, CardBody, Spinner } from "reactstrap";

export default function Loading() {
  return (
    <Card>
      <CardBody className="mx-auto">
        <Spinner className="mx-auto" />
      </CardBody>
    </Card>
  );
}

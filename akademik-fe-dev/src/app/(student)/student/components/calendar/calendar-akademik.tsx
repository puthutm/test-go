import Image from "next/image";
import { Card, CardBody } from "reactstrap";

export default function CalendarAcademicView() {
  return (
    <Card className="mt-4" style={{ borderRadius: "8px" }}>
      <CardBody className="d-flex justify-content-center align-items-center py-3 mb-0 gap-4">
        <Image
          src={"/kalender-akademik.jpg"}
          alt="Kalender Akademik"
          width={800}
          height={700}
        />
      </CardBody>
    </Card>
  );
}

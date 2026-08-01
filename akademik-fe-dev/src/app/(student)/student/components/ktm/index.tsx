import { FileDownloadIcon } from "@/components/icons/file-download";
import { ForwardToInboxIcon } from "@/components/icons/forward-to-inbox";
import { Card, CardBody, CardHeader } from "reactstrap";

import KtmFront from "@/assets/images/ktm-front.png";
import KtmBack from "@/assets/images/ktm-back.png";
import Image from "next/image";

export default function KTMView() {
  return (
    <Card className="mt-4" style={{ borderRadius: "8px" }}>
      <CardHeader
        className="border border-bottom-0 pb-0"
        style={{ borderRadius: "8px" }}
      >
        <div className="d-flex justify-content-between align-items-center pb-3 border-2 border-bottom ">
          <h2
            className={`card-title fw-medium fs-5 mb-0`}
            style={{ color: "#495057" }}
          >
            Kartu Tanda Mahasiswa
          </h2>
          <div className="d-flex gap-3">
            <button
              // onClick={handleSendEmail}
              className="btn d-flex align-items-center gap-2 text-primary"
              style={{
                whiteSpace: "nowrap",
                border: "1px solid #10487A",
                backgroundColor: "transparent",
              }}
            >
              <ForwardToInboxIcon /> Kirim ke Email
            </button>
            <button
              // onClick={handleDownloadKHS}
              className="btn d-flex align-items-center gap-2 text-primary"
              style={{
                whiteSpace: "nowrap",
                border: "1px solid #10487A",
                backgroundColor: "transparent",
              }}
            >
              <FileDownloadIcon /> Download
            </button>
          </div>
        </div>
      </CardHeader>
      <CardBody className="d-flex justify-content-center align-items-center py-3 mb-0 gap-4">
        <Image src={KtmFront} alt="ktm_front" width={515} />
        <Image src={KtmBack} alt="ktm_back" width={515} />
      </CardBody>
    </Card>
  );
}

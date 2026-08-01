import { Card, CardBody, Col, Row } from "reactstrap";

import { SideTab } from "../side-tabs";
import { PersonIcon } from "@/components/icons/person";
import { ArticleIcon } from "@/components/icons/article";
import { FormBiodata } from "./form-biodata";
import { FolderIcon } from "@/components/icons/folder";
import { GroupIcon } from "@/components/icons/group";
import { DescriptionIcon } from "@/components/icons/description";
import { CreditCardIcon } from "@/components/icons/credit-card";
import { MapIcon } from "@/components/icons/map";
import { SchoolIcon } from "@/components/icons/school";
import { FormOriginalEducation } from "./form-original-education";
import { FormInformation } from "./form-information";
import { FormAddress } from "./form-address";
import { FormDocument } from "./form-document";
import { FormBankAccount } from "./form-account-bank";
import { FormCompleteness } from "./form-completeness";
import { FormParent } from "./form-parent";
import { getBiodataStudent } from "@/services/api/students/biodata/biodata/get-biodatas";
import { getInformationStudent } from "@/services/api/students/biodata/information/get-informations";
import { getAddressStudent } from "@/services/api/students/biodata/address/get-address";

export default async function TabBiodata({
  searchParams,
}: {
  searchParams: { [key: string]: string | undefined };
}) {
  const stateParam = searchParams?.state;

  const [biodataResponse, informationResponse, addressResponse] =
    await Promise.all([
      getBiodataStudent(),
      getInformationStudent(),
      getAddressStudent(),
    ]);

  const tabs = [
    {
      label: "Biodata",
      key: "biodata",
      icon: (
        <PersonIcon
          color={stateParam === "biodata" || !stateParam ? "white" : "#495057"}
        />
      ),
    },
    {
      label: "Informasi",
      key: "information",
      icon: (
        <ArticleIcon
          color={stateParam === "information" ? "white" : "#495057"}
        />
      ),
    },
    {
      label: "Kelengkapan",
      key: "completness",
      icon: (
        <FolderIcon
          color={stateParam === "completness" ? "white" : "#495057"}
        />
      ),
    },
    {
      label: "Alamat",
      key: "address",
      icon: <MapIcon color={stateParam === "address" ? "white" : "#495057"} />,
    },
    {
      label: "Data Orang Tua",
      key: "family",
      icon: <GroupIcon color={stateParam === "family" ? "white" : "#495057"} />,
    },
    {
      label: "Dokumen",
      key: "document",
      icon: (
        <DescriptionIcon
          color={stateParam === "document" ? "white" : "#495057"}
        />
      ),
    },
    {
      label: "Rekening",
      key: "account_bank",
      icon: (
        <CreditCardIcon
          color={stateParam === "account_bank" ? "white" : "#495057"}
        />
      ),
    },
    {
      label: "Pendidikan Asal",
      key: "school",
      icon: (
        <SchoolIcon color={stateParam === "school" ? "white" : "#495057"} />
      ),
    },
  ];

  return (
    <Row className="pt-4">
      <Col lg={3}>
        <Card style={{ position: "sticky", top: 80, borderRadius: "8px" }}>
          <CardBody>
            <SideTab tabs={tabs} />
          </CardBody>
        </Card>
      </Col>
      <Col lg={9}>
        <Card style={{ borderRadius: "8px" }}>
          <CardBody>
            {stateParam === "biodata" || !stateParam ? (
              <FormBiodata biodata={biodataResponse} />
            ) : null}
            {stateParam === "information" ? (
              <FormInformation information={informationResponse} />
            ) : null}
            {stateParam === "completness" ? <FormCompleteness /> : null}
            {stateParam === "address" ? (
              <FormAddress address={addressResponse} />
            ) : null}
            {stateParam === "family" ? <FormParent /> : null}
            {stateParam === "document" ? <FormDocument /> : null}
            {stateParam === "account_bank" ? <FormBankAccount /> : null}
            {stateParam === "school" ? <FormOriginalEducation /> : null}
          </CardBody>
        </Card>
      </Col>
    </Row>
  );
}

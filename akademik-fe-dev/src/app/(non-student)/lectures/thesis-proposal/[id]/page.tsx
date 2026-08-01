import React from "react";
import TAProposalTabContent from "./_components/tabs/ta-proposal";
import ThesisTabContent from "./_components/tabs/thesis";
import ConsultationTabContent from "./_components/tabs/consultation";
import RequirementTabContent from "./_components/tabs/requirement";

export default function ThesisProposalPage({
  searchParams,
}: {
  searchParams: { tabs: string };
}) {
  const { tabs } = searchParams;

  switch (tabs) {
    case "proposal":
      return <TAProposalTabContent />;
    case "thesis":
      return <ThesisTabContent />;
    case "consultation":
      return <ConsultationTabContent />;
    case "requirements":
      return <RequirementTabContent />;
    default:
      return (
        <div>
          <h1>Thesis Proposal Page</h1>
          <p>tabs: {searchParams.tabs}</p>
        </div>
      );
  }
}

// ReportTable.js
import React, { useState, useEffect } from "react";

const ReportTable = () => {
  const [htmlTable, setHtmlTable] = useState("");

  useEffect(() => {
    fetch("http://localhost:8080/api/reports/html", {
      method: "GET",
    })
      .then((res) => res.text())
      .then((data) => {
        setHtmlTable(data);
      })
      .catch((error) => console.error("Error fetching HTML reports:", error.message));
  }, []);

  return (
    <div dangerouslySetInnerHTML={{ __html: htmlTable }} />
  );
};

export default ReportTable;

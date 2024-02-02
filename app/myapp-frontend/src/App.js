import React, { useState, useEffect } from "react";
import "bootstrap/dist/css/bootstrap.min.css";
import ReportModal from "./ReportModal";
import ReportTable from "./ReportTable";

const App = () => {
  const [reports, setReports] = useState([]);
  const [modalShow, setModalShow] = useState(false);
  const [selectedReport, setSelectedReport] = useState(null);

  useEffect(() => {
    fetch("http://localhost:8080/api/reports", {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
      },
    })
      .then((res) => res.json())
      .then((data) => {
        console.log(data);
        setReports(data);
      })
      .catch((error) => console.error("Error fetching reports:", error.message));
  }, []);

  const handleReportClick = (report) => {
    setSelectedReport(report);
    setModalShow(true);
  };

  return (
    <div className="container mt-5">
      <nav className="navbar navbar-light bg-success">
        <div className="container-fluid">
          <span className="navbar-brand mb-0 h1 text-white">Azen mdl</span>
        </div>
      </nav>

      <h1 className="mt-3"> Reportes X </h1>

      <button className="btn btn-success" onClick={() => setModalShow(true)}>
        Reportes Azen
      </button>

      <div className="mt-3">
        <ReportTable reports={reports} onReportClick={handleReportClick} />
      </div>

      <ReportModal
        reports={reports}
        show={modalShow}
        handleClose={() => setModalShow(false)}
        handleReportClick={handleReportClick}
        selectedReport={selectedReport}
      />
    </div>
  );
};

export default App;

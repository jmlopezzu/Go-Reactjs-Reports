import React, { useState } from "react";

const ReportModal = ({ reports, show, handleClose, handleReportClick }) => {
  const [selectedReport, setSelectedReport] = useState(null);

  const handleItemClick = (report) => {
    setSelectedReport(report);
    handleReportClick(report); // Pass the selected report to the parent component if needed
  };

  return (
    <div className={`modal ${show ? "show" : ""}`} tabIndex="-1" role="dialog" style={{ display: show ? "block" : "none" }}>
      <div className="modal-dialog" role="document">
        <div className="modal-content">
          <div className="modal-header bg-success text-white">
            <h5 className="modal-title">Azen + Data</h5>
            <button type="button" className="close" data-dismiss="modal" aria-label="Close" onClick={handleClose}>
              <span aria-hidden="true">&times;</span>
            </button>
          </div>
          <div className="modal-body">
          
          
          {selectedReport && (
              <div>
                <h1>Reportes Especificos </h1>
                <br></br>
                <h3>ID Cliente: {selectedReport.id}</h3>
                <br></br>
                <p>Detalle : {selectedReport.content}</p>
              </div>
            )}



            {reports?.length > 0 ? (
              <ul className="list-group">
                {reports.map(({ id, content }) => (
                  <li key={id} className="list-group-item">
                    <a href="#!" onClick={() => handleItemClick({ id, content })}>
                      {id}
                    </a>
                  </li>
                ))}
              </ul>
            ) : (
              <p>No hay reportes XD </p>
            )}

            {/* Display detailed information if a report is selected */}
            
          </div>
        </div>
      </div>
    </div>
  );
};

export default ReportModal;

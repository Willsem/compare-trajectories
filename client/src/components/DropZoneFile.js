import React, { useMemo, useState } from 'react';
import { useDropzone } from 'react-dropzone';

const baseStyle = {
  flex: 1,
  display: 'flex',
  flexDirection: 'column',
  alignItems: 'center',
  padding: '20px',
  height: '100%',
  borderWidth: 2,
  justifyContent: 'center',
  borderRadius: 10,
  borderColor: '#eeeeee',
  borderStyle: 'dashed',
  backgroundColor: '#fafafa',
  color: '#bdbdbd',
  outline: 'none',
  transition: 'border .24s ease-in-out',
};

const activeStyle = {
  borderColor: '#2196f3',
};

const acceptStyle = {
  borderColor: '#00e676',
};

const rejectStyle = {
  borderColor: '#ff1744',
};

let lastFile = '';

function DropZoneFile({ loadFileCallback }) {
  const [, setFile] = useState(0);

  const {
    getRootProps,
    getInputProps,
    isDragActive,
    isDragAccept,
    isDragReject,
    acceptedFiles,
  } = useDropzone({accept: 'application/json'});

  const style = useMemo(() => ({
    ...baseStyle,
    ...(isDragActive ? activeStyle : {}),
    ...(isDragAccept ? acceptStyle : {}),
    ...(isDragReject ? rejectStyle : {})
  }), [
    isDragActive,
    isDragReject,
    isDragAccept
  ]);

  acceptedFiles.forEach(file => {
    if (file.name !== lastFile) {
      lastFile = file.name;
      let reader = new FileReader();
      reader.readAsText(file);

      reader.onload = function() {
        const parsedFile = JSON.parse(reader.result);
        setFile(parsedFile);
        loadFileCallback(parsedFile);
      };
    }
  });

  return (
    <div className="container">
      <div {...getRootProps({style})}>
        <input {...getInputProps()} />
        <p>Drag 'n' drop json file here, or click to select it</p>
      </div>
    </div>
  );
}

export default DropZoneFile;

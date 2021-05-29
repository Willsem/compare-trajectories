import React, { useMemo } from 'react';
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

let lastFile = {};

function DropZoneFile({ loadFileCallback, fieldName }) {
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
    if (file.name !== lastFile[fieldName]) {
      lastFile[fieldName] = file.name;
      let reader = new FileReader();
      reader.readAsText(file);

      reader.onload = function() {
        const parsedFile = JSON.parse(reader.result);
        loadFileCallback(parsedFile);
      };
    }
  });

  return (
    <div className="container">
      <div {...getRootProps({style})}>
        <input {...getInputProps()} />
        <p>Перетащите сюда json-файл или нажмите, чтобы загрузить его</p>
      </div>
    </div>
  );
}

export default DropZoneFile;

import React from 'react';
import NavBar2 from '../components/navbar2';
import Card from '../components/card';
import ReactEcharts from 'echarts-for-react';
import '../components/styles/TaskManager.css';
import '../components/styles/input.css';
// Función para generar un color aleatorio en formato hexadecimal
const getRandomColor = () => {
  const letters = '0123456789ABCDEF';
  let color = '#';
  for (let i = 0; i < 6; i++) {
    color += letters[Math.floor(Math.random() * 16)];
  }
  return color;
};

// Datos para los gráficos
const ramData = {
  used: 80,
  free: 20
};

const processorData = {
  used: 35,
  free: 65
};

export default function TaskManager() {

  const ramm = {
    tooltip: {
      trigger: 'item'
    },
    legend: {
      top: '5%',
      left: 'center',
      textStyle: {
        color: '#ffffff' // Cambia el color de las etiquetas de la leyenda
      }
    },
    series: [
      {
        name: 'RAM',
        type: 'pie',
        radius: ['40%', '70%'],
        avoidLabelOverlap: false,
        padAngle: 5,
        itemStyle: {
          borderRadius: 10,
          color: function (params) {
            // Genera un color aleatorio para cada sector
            return getRandomColor();
          }
        },
        label: {
          show: false, // Puedes cambiar esto a true si quieres mostrar las etiquetas
          position: 'outside',
          textStyle: {
            color: 'white' // Cambia el color de las etiquetas de los datos
          }
        },
        emphasis: {
          label: {
            show: false,
            fontSize: 20,
            fontWeight: 'bold',
            color: '#ff0000' // Cambia el color de la etiqueta en el énfasis
          }
        },
        labelLine: {
          show: true,
          lineStyle: {
            color: '#ffffff' // Cambia el color de las líneas de las etiquetas
          }
        },
        data: [
          { value: ramData.used, name: 'Used' },
          { value: ramData.free, name: 'Free' },
        ]
      }
    ]
  };

  const processor = {
    tooltip: {
      trigger: 'item'
    },
    legend: {
      top: '5%',
      left: 'center',
      textStyle: {
        color: '#ffffff' // Cambia el color de las etiquetas de la leyenda
      }
    },
    series: [
      {
        name: 'Processor',
        type: 'pie',
        radius: ['40%', '70%'],
        avoidLabelOverlap: false,
        padAngle: 5,
        itemStyle: {
          borderRadius: 10,
          color: function (params) {
            // Genera un color aleatorio para cada sector
            return getRandomColor();
          }
        },
        label: {
          show: false, // Puedes cambiar esto a true si quieres mostrar las etiquetas
          position: 'outside',
          textStyle: {
            color: 'white' // Cambia el color de las etiquetas de los datos
          }
        },
        emphasis: {
          label: {
            show: false,
            fontSize: 20,
            fontWeight: 'bold',
            color: '#ff0000' // Cambia el color de la etiqueta en el énfasis
          }
        },
        labelLine: {
          show: true,
          lineStyle: {
            color: '#ffffff' // Cambia el color de las líneas de las etiquetas
          }
        },
        data: [
          { value: processorData.used, name: 'Used' },
          { value: processorData.free, name: 'Free' },
        ]
      }
    ]
  };

  const data = [
    { PID: 1234, Name: 'Process 1', State: 'Running' },
    { PID: 5678, Name: 'Process 2', State: 'Sleeping' },
    { PID: 9101, Name: 'Process 3', State: 'Stopped' }
  ];


  const handleFileUpload = (e) => {
    // Manejar la carga de archivos aquí
  };

  const handleMessageInputChange = (e) => {
    // Manejar el cambio en el input de mensaje aquí
  };

  const handleSendMessageClick = () => {
    // Manejar el clic en el botón de enviar aquí
  };


  return (
    <div>
      <NavBar2 />
      <div style={{ display: 'flex', flexDirection: 'column', alignItems: 'center', marginTop: '20px' }}>
        <Card>
          <div className="card-title">Graphs</div>
          <div style={{ display: 'flex', justifyContent: 'space-between', width: '100%' }}>
            <div style={{ flex: 1, margin: '0 10px' }}>
              <h3>Ram</h3>
              <ReactEcharts option={ramm} />
            </div>
            <div style={{ flex: 1, margin: '0 10px' }}>
              <h3>Processor</h3>
              <ReactEcharts option={processor} />
            </div>
          </div>
        </Card>
        <Card>
          <div className="card-title">Tasks</div>
          <div style={{ display: 'flex', flexDirection: 'column', alignItems: 'center' }}>
            <div className="messageBox">
              <div className="fileUploadWrapper">
                <label htmlFor="file">
                  <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 337 337">
                    <circle strokeWidth="20" stroke="#6c6c6c" fill="none" r="158.5" cy="168.5" cx="168.5"></circle>
                    <path strokeLinecap="round" strokeWidth="25" stroke="#6c6c6c" d="M167.759 79V259"></path>
                    <path strokeLinecap="round" strokeWidth="25" stroke="#6c6c6c" d="M79 167.138H259"></path>
                  </svg>
                  <span className="tooltip">Add a Process</span>
                </label>
                <input type="file" id="file" name="file" onChange={handleFileUpload} />
              </div>
              <div className="inputWithButtons">
                <input required placeholder="PID..." type="text" id="messageInput" onChange={handleMessageInputChange} />
              </div>
              <button id="sendButton" onClick={handleSendMessageClick}>
                <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 50 50" width="50px" height="50px">
                  <path fill="red"d="M 21 2 C 19.354545 2 18 3.3545455 18 5 L 18 7 L 10.154297 7 A 1.0001 1.0001 0 0 0 9.984375 6.9863281 A 1.0001 1.0001 0 0 0 9.8398438 7 L 8 7 A 1.0001 1.0001 0 1 0 8 9 L 9 9 L 9 45 C 9 46.645455 10.354545 48 12 48 L 38 48 C 39.645455 48 41 46.645455 41 45 L 41 9 L 42 9 A 1.0001 1.0001 0 1 0 42 7 L 40.167969 7 A 1.0001 1.0001 0 0 0 39.841797 7 L 32 7 L 32 5 C 32 3.3545455 30.645455 2 29 2 L 21 2 z M 21 4 L 29 4 C 29.554545 4 30 4.4454545 30 5 L 30 7 L 20 7 L 20 5 C 20 4.4454545 20.445455 4 21 4 z M 11 9 L 18.832031 9 A 1.0001 1.0001 0 0 0 19.158203 9 L 30.832031 9 A 1.0001 1.0001 0 0 0 31.158203 9 L 39 9 L 39 45 C 39 45.554545 38.554545 46 38 46 L 12 46 C 11.445455 46 11 45.554545 11 45 L 11 9 z M 18.984375 13.986328 A 1.0001 1.0001 0 0 0 18 15 L 18 40 A 1.0001 1.0001 0 1 0 20 40 L 20 15 A 1.0001 1.0001 0 0 0 18.984375 13.986328 z M 24.984375 13.986328 A 1.0001 1.0001 0 0 0 24 15 L 24 40 A 1.0001 1.0001 0 1 0 26 40 L 26 15 A 1.0001 1.0001 0 0 0 24.984375 13.986328 z M 30.984375 13.986328 A 1.0001 1.0001 0 0 0 30 15 L 30 40 A 1.0001 1.0001 0 1 0 32 40 L 32 15 A 1.0001 1.0001 0 0 0 30.984375 13.986328 z" />
                </svg>
              </button>
            </div>
          </div>
          <div className="table-container">
            <table className="cyberpunk-table">
              <caption className="table-caption">Process List</caption>
              <thead>
                <tr>
                  <th>PID</th>
                  <th>Name</th>
                  <th>State</th>
                </tr>
              </thead>
              <tbody>
                {data.map((item, index) => (
                  <tr key={index}>
                    <td>{item.PID}</td>
                    <td>{item.Name}</td>
                    <td>{item.State}</td>
                  </tr>
                ))}
              </tbody>
            </table>
          </div>
        </Card>
      </div>
    </div>
  );
}

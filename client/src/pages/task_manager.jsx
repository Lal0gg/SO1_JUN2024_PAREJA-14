import NavBar2 from '../components/navbar2';
import Card from '../components/card';
import ReactEcharts from 'echarts-for-react';

// Función para generar un color aleatorio en formato hexadecimal
const getRandomColor = () => {
  const letters = '0123456789ABCDEF';
  let color = '#';
  for (let i = 0; i < 6; i++) {
    color += letters[Math.floor(Math.random() * 16)];
  }
  return color;
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
          { value: 80, name: 'Used' },
          { value: 20, name: 'Free' },
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
          { value: 80, name: 'Used' },
          { value: 20, name: 'Free' },
        ]
      }
    ]
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
          <div>
            {/* Aquí puedes añadir contenido relacionado con las tareas */}
          </div>
        </Card>
      </div>
    </div>
  );
}

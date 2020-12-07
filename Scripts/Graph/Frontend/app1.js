document.getElementById("getUpdates").addEventListener("click", function(){
    axios.get('/simulate')
})

const pusher = new Pusher('7befe6ab035a03a2ada9', {
    cluster: 'us2',
    encrypted: true
});

const channel = pusher.subscribe('randomNum');


channel.bind('addNumber', data => {

    if (newLineChart.data.datasets[0].data.length > 0) 
    {
            newLineChart.data.datasets[0].data=[]
            newLineChart.data.labels=[]
    }
    data.forEach(function (value, i){
        
        newLineChart.data.labels.push(i+1);
        newLineChart.data.datasets[0].data.push(value);
    });
    newLineChart.update();
});

function renderChart(userVisitsData) {
  var ctx = document.getElementById("realtimeChart").getContext("2d");

  var options = {};

  newLineChart = new Chart(ctx, {
    type: "bar",
    data: userVisitsData,
    options: options
  });
}

var chartConfig = {
  labels: [],
  datasets: [
     {
        label: "Magnitud de los números pseudoaleatorios",
        fill: false,
        lineTension: 0.1,
        backgroundColor: "rgba(75,192,192,0.4)",
        borderColor: "rgba(75,192,192,1)",
        borderCapStyle: 'butt',
        borderDash: [],
        borderDashOffset: 0.0,
        borderJoinStyle: 'miter',
        pointBorderColor: "rgba(75,192,192,1)",
        pointBackgroundColor: "#fff",
        pointBorderWidth: 1,
        pointHoverRadius: 5,
        pointHoverBackgroundColor: "rgba(75,192,192,1)",
        pointHoverBorderColor: "rgba(220,220,220,1)",
        pointHoverBorderWidth: 2,
        pointRadius: 1,
        pointHitRadius: 10,
        data: [],
        spanGaps: false,
     }
  ]
};

renderChart(chartConfig)
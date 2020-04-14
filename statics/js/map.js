var attackData = [];
var series = [];

function getData() {
    var attackPostion = [];
    $.ajaxSettings.async = false;
    $.get("static/json/IOC", function (response) {
        response.data.forEach(element => {
            attackPostion.push([element.sourcePosition, element.desPosition]);
        });
    })
    $.ajaxSettings.async = true;
    return attackPostion;
}

var dser = {
    type: 'effectScatter',
    coordinateSystem: 'geo',
    zlevel: 3,
    rippleEffect: {
        brushType: 'stroke'
    },
    label: {
        normal: {
            show: true,
            position: 'left',
            fontSize: 15,
            formatter: '{b}'
        }
    },
    itemStyle: {
        normal: {
            color: '#ff0000'
        }
    }
}
attackData = getData();
series = {
    type: 'lines3D',
    coordinateSystem: 'globe',
    effect: {
        show: true,
        period: 2.5, //速度
        trailLength: 0.5 //尾部阴影          
    },
    blendMode: 'lighter',
    lineStyle: { //航线的视图效果
        //  color: '#000',
        width: 3,
        opacity: 1
    },
    data: attackData // 特效的起始、终点位置，一个二维数组，相当于coords: convertData(item[1])
}
//  });
/**
 * 地球皮肤
 */
console.log(series);

var canvas = document.createElement('canvas');

var myChart = echarts.init(canvas, null, {
    width: 2048,
    height: 1024
});
myChart.setOption({
    backgroundColor: "#13334c",
    title: {
        show: true
    },
    geo: {
        type: 'map',
        map: 'world',
        left: 0,
        top: 0,
        right: 0,
        bottom: 0,
        boundingCoords: [
            [-180, 90],
            [180, -90]
        ],
        zoom: 0,
        roam: false,
        itemStyle: {
            borderColor: "#f6f6e9",
            normal: {
                areaColor: '#005792',
                borderColor: '#f6f6e9'
            },
            emphasis: {
                areaColor: '#ff6768' //当鼠标移入时的显示
            }
        },
        label: {
            fontSize: 20
        }
    },
    series: dser
});

/**
 * 3D地球
 */
var option = {
    backgroundColor: '#17223b', //canvas的背景颜色
    globe: {
        baseTexture: myChart,
        top: 'top',
        left: 'right',
        displacementScale: 0,
        environment: 'none',
        shading: 'color',
        viewControl: {
            distance: 180,
            zoomSensitivity: 0,
            // maxDistance :240, //最大的值 （默认400）
            // minDistance :240, //是距离 最小值 （默认40） 与最大值相等时 则不能够放大与缩小
            autoRotate: true,
        },
    },
    series: series,
};

echarts.init(document.getElementById("map")).setOption(option, true);
(function(){
      function getDaysAgoDateHour(days){
        var _daysAgo = new Date();
        var zPad = function(str, order){
            if (typeof order == 'undefined') order = 2;
            str = String(str);
            for (var i = 0; i<order; i++) {
                str = '0' + str;
            }
            return str.slice((-1)*order);
        };
        _daysAgo.setDate(_daysAgo.getDate() - parseInt(days));
        return _daysAgo.getFullYear() + zPad(_daysAgo.getMonth() + 1) + zPad(_daysAgo.getDate()) + zPad(_daysAgo.getHours());
      }
      google.load("visualization", "1", {packages:["corechart"]});
      google.setOnLoadCallback(drawChart);
      function drawChart() {
        var options = {
          title: 'OCR Requests',
          hAxis: {title: 'DATE',  titleTextStyle: {color: '#333'}},
          vAxis: {
            0 : {minValue: 0},
            1 : {minValue: 0},
            maxValue : 100
          },
          seriesType : "area",
          isStacked : true,
          series : {
            0 : {targetAxisIndex : 0, color : '#dc3912'},
            1 : {targetAxisIndex : 0, color : '#3366cc'},
            2 : {targetAxisIndex : 1 , isStacked : false , type:'line'}
          }
        };
        var data = new google.visualization.DataTable();
        var charData = [["Date","Failure","Success"]];
        data.addColumn('string','date');
        data.addColumn('number','Failure');
        data.addColumn('number','Success');
        data.addColumn('number','Success Rate(%)');
        var url = '/ocr/summary/' + getDaysAgoDateHour(12);
        $.get(url,{},
          function(summary){
            console.log(summary);
            $.each(summary, function(){
              var index = this.Month+'/'+this.Date+'/'+this.Hour+':00';
              var rate = Math.floor(this.Success * 100 / (this.Success + this.Failure));
              data.addRow([index, this.Failure, this.Success, rate]);
            });
            var chart = new google.visualization.ComboChart(document.getElementById('ocr-summary-chart'));
            chart.draw(data, options);
          }
        );
      }
})();

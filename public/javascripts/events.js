$(function(){
  $('input.delelte-record').on('click',function(){
    var target = $(this).attr('target');
    var doneCallback = function(res){
      $('tr#' + target).fadeOut(function(){
        $(this).hide().remove();
      });
    };
    var data = { target : target };
    var url  = '/ocr/delete';
    var message = 'DELETE this?\n' + target;
    if(window.confirm(message)){
      $.post(
        url,
        data,
        doneCallback
      );
    }
  });
  setInterval(function(){
    if (localStorage.getItem('autoReload') == 'true') {
      location.reload();
    }
  },20*1000);
  $('input#option-auto-reload').on('change',function(){
    localStorage.setItem('autoReload', $(this).is(':checked'));
  });
  if (localStorage.getItem('autoReload') == 'true') {
    $('input#option-auto-reload').attr('checked',true);
  }
});

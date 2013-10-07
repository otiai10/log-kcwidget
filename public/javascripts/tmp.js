$(function(){
  $('a.created-time').each(function(){
    var s = $(this).html();
    var d = new Date(parseInt(s));
    $(this).html(d.toLocaleString());
  });
  $('.proxy').each(function(){
    var uri = $(this).html();
    var img = new Image();
    img.src = uri;
    $(this).after(img.outerHTML);
    $(this).remove();
  });
});

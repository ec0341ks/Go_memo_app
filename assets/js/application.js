require("expose-loader?$!expose-loader?jQuery!jquery");
require("bootstrap/dist/js/bootstrap.bundle.js");



$('.delete-button').on('click', function (e) {
  e.preventDefault();
  var nameid = $(this).closest(".profile").children(".id").text();
  var checked = confirm("本当に削除しますか？")
  if (checked == true) {
    $(this).closest(".profile").remove();
    $.ajax({
      type: 'DELETE',
      url: '/boards/' + nameid,
      dataType: 'json'
    })
    alert("削除しました")
  }
  else {
    alert("削除されませんでした")
  }
})
require("expose-loader?$!expose-loader?jQuery!jquery");
require("bootstrap/dist/js/bootstrap.bundle.js");



$('.profile').on('click', function () {
  $(this).remove();
  console.log("発火！")
})
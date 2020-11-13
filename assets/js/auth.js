
var isLogin = false;

function showLoginCnt() {
    isLogin = true
    $("#login_container").css('visibility', 'visible');
    $("#login_container").css('position', 'relative');
    $("#login_container").css('height', '100%');

    $("#register_container").css('visibility', 'hidden');
    $("#register_container").css('position', 'absolute');
    $("#register_container").css('height', '0');
}

function hideLoginCnt() {
    isLogin = false
    $("#login_container").css('visibility', 'hidden');
    $("#login_container").css('position', 'absolute');
    $("#login_container").css('height', '0');

    $("#register_container").css('visibility', 'visible');
    $("#register_container").css('position', 'relative');
    $("#register_container").css('height', '100%');
}
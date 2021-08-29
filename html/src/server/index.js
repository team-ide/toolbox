import http from '@/server/http';
import tool from "@/tool/index.js";
let server = {
    validate(param) {
        return new Promise((resolve, reject) => {
            http.post('server/validate', param, {
                headers: {}
            }).then((res) => {
                if (res.value) {

                    if (tool.TIMESTAMP != null) {
                        if (tool.TIMESTAMP != res.value) {
                            tool.alert('服务器已重启，请刷新页面！').then(() => {

                            });
                            return;
                        }
                    } else {
                        tool.TIMESTAMP = res.value;
                    }
                }
                resolve(true);
            }).catch(e => {
                resolve(false);
            });
        });
    },
    wait(wait_check_count) {
        wait_check_count = wait_check_count || 0;
        return new Promise((resolve, reject) => {
            server.validate().then(res => {
                if (res) {
                    resolve && resolve(wait_check_count);
                } else {
                    wait_check_count++;
                    let second = 3;
                    if (wait_check_count >= 3) {
                        second = 5;
                    }
                    if (wait_check_count >= 5) {
                        second = 10;
                    }
                    if (wait_check_count >= 10) {
                        second = 30;
                    }
                    if (wait_check_count >= 20) {
                        second = 60;
                        return;
                    }
                    window.setTimeout(function () {
                        server.wait(wait_check_count).then(resolve);
                    }, second * 1000);
                }
            });
        });
    },
    waitStop(wait_check_count) {
        wait_check_count = wait_check_count || 0;
        return new Promise((resolve, reject) => {
            server.validate().then(res => {
                if (!res) {
                    resolve && resolve(wait_check_count);
                } else {
                    wait_check_count++;
                    let second = 3;
                    if (wait_check_count >= 3) {
                        second = 5;
                    }
                    if (wait_check_count >= 5) {
                        second = 10;
                    }
                    if (wait_check_count >= 10) {
                        second = 30;
                    }
                    if (wait_check_count >= 20) {
                        second = 60;
                        return;
                    }
                    window.setTimeout(function () {
                        server.waitStop(wait_check_count).then(resolve);
                    }, second * 1000);
                }
            });
        });
    },
    open(param) {
        return http.post('server/open', param, {
            headers: {
            }
        });
    },
    close(param) {
        return http.post('server/close', param, {
            headers: {
            }
        });
    },
    session(param) {
        return http.post('server/session', param, {
            headers: {}
        });
    },
};


export default server;
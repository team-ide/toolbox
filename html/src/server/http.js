import axios from 'axios';
import tool from "@/tool/index.js";
import source from "@/source/index.js";

let axiosConfig = {};
axiosConfig.timeout = 1000 * 60 * 10;
axiosConfig.baseURL = "";
axiosConfig.withCredentials = true; // 请求带上cookie
let axiosInstance = axios.create(axiosConfig);
// 添加请求拦截器
axiosInstance.interceptors.request.use(function (config) {
	// 在发送请求之前做些什么
	config = config || {};
	config.headers = config.headers || {};
	return config;
}, function (error) {
	// 对请求错误做些什么
	return Promise.reject(error);
});

// 添加响应拦截器
axiosInstance.interceptors.response.use(function (response) {
	if (response.config && response.config.responseType == 'blob') {

		let blob = new Blob([response.data]);
		let downloadElement = document.createElement('a')
		let href = window.URL.createObjectURL(blob); //创建下载的链接
		downloadElement.href = href;
		downloadElement.download = response.headers['download-file-name']; //下载后文件名
		document.body.appendChild(downloadElement);
		downloadElement.click(); //点击下载
		document.body.removeChild(downloadElement); //下载完成移除元素
		window.URL.revokeObjectURL(href); //释放blob对象

		return response.data;
	}
	if (response.data) {
		const code = '' + response.data.code;
		switch (code) {
			case "0":
				return response.data;
			case "10001":
				tool.error('客户端会话失效，请刷新浏览器！');
				return response.data;
			case "10002":
				tool.error('客户端会话失效，请刷新浏览器！');
				return response.data;
			case "10101":
				tool.error('暂无登录信息，请先登录！');
				source.login_user = null;
				return response.data;
			case "10102":
				tool.error('暂无权限执行此次操作！');
				return response.data;
		}
	}

	// 对响应数据做点什么
	return response.data;
}, function (error) {
	// 对响应错误做点什么
	return Promise.reject(error);
});

export default axiosInstance;
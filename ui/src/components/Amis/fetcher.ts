import {FetcherConfig} from "amis-core/lib/factory";
import {fetcherResult} from "amis-core/lib/types";
import axios from "axios";
import {Message} from "@arco-design/web-react";

export const fetcher = ({url, method = 'get', data, config}: FetcherConfig): Promise<fetcherResult> => {
    const token = localStorage.getItem('token') || '';

    const ajax = axios.create({
        baseURL: '/',
        headers: {
            ...config?.headers,
            Authorization: token ? `Bearer ${token}` : ''
        }
    });

    // 请求发送之前的拦截
    ajax.interceptors.response.use(
        response => response, // 请求成功
        error => {
            if (error.response && error.response.status === 401) {
                // 如果是401，跳转到登录页面
                window.location.href = '/#/login';
            }
            if (error.response && error.response.status === 512) {
                var cluster = error.response.data.msg;
                Message.error(`集群【${cluster}】尚未连接成功，请先连接，或选择其他集群`)
                window.location.href = '/#/cluster/cluster_all';
            }
            return Promise.reject(error); // 继续处理其他错误
        }
    );

    switch (method.toLowerCase()) {
        case 'get':
            return ajax.get(url, config);
        case 'post':
            return ajax.post(url, data || null, config);
        default:
            return ajax.post(url, data || null, config);
    }
};

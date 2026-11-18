
const API_URL = import.meta.env.VITE_API_URL || 'http://localhost:8080';

export const api = await fetch({
	baseURL: API_URL,
	headers: {
		'Content-Type': 'application/json'
	}
});

// Interceptor: dodaj token
api.interceptors.request.use((config) => {
	const token = localStorage.getItem('token');
	if (token) {
		config.headers.Authorization = `Bearer ${token}`;
	}
	return config;
});

// Interceptor: 401 → redirect do loginu
api.interceptors.response.use(
	(res) => res,
	(err) => {
		if (err.response?.status === 401) {
			localStorage.removeItem('token');
			window.location.href = '/login';
		}
		return Promise.reject(err);
	}
);


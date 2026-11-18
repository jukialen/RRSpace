import { publicPaths, staticPaths } from 'myConstants/links';

export const isPublicPath = (pathname: string) => {
	return publicPaths.some(
		(publicPath) => pathname === publicPath || pathname.startsWith(publicPath + '/')
	);
};

export const isStaticPath = (pathname: string) => {
	return staticPaths.some(
		(staticPath) => pathname === staticPath || pathname.startsWith(staticPath + '/')
	);
};

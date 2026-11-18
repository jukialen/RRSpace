import { fail } from '@sveltejs/kit';
import * as Yup from 'yup';
import type { Actions, PageServerLoad } from './$types';
import { m } from '$lib/paraglide/messages';


import { SchemaValidation } from '$lib/helpers/schemaValidation/schemaValidation';

export const load: PageServerLoad = async ({ cookies, parent, depends }) => {
	// This will ensure that `invalidate('supabase:auth')` works on this page
	depends('backend:auth');

	const data = await parent(); // Inherits the user from the layout

	const firstLogin = cookies.get('first-login');

	return { ...data, getFirstLogin: firstLogin };
};

export const actions = {
	default: async ({ request }) => {
		try {
			const data = await request.formData();
			const email = data.get('email') as string;
			const password = data.get('password') as string;

			const schemaValidation = Yup.object({
				email: SchemaValidation().email,
				password: SchemaValidation().password
			});

			await schemaValidation.validate({ email, password }, { abortEarly: false });

			const { data, error }: { data: string; error: string } = await fetch('/auth/login', {
				method: 'POST',
				body: JSON.stringify({ email, password })
			});


			if (error) {
				const unVerifedResponse =
					error?.includes('Email not confirmed') ||
					error.includes('AuthApiError: Email not confirmed') ||
					error?.includes('Invalid login credentials');

				return fail(400, {
					errors: { email: '', password: '' },
					message: unVerifedResponse ? m.NavForm_unVerified() : m.NavForm_wrongLoginData(),
					success: false
				});
			}

			if (!data) {
				return fail(400, {
					errors: { email: '', password: '' },
					message: m.form_communication_server_error(),
					success: false
				});
			}

			return {
				errors: { email: '', password: '' },
				success: true,
				message: m.NavForm_statusLogin(),
			};
		} catch (e) {
			if (e instanceof Yup.ValidationError) {
				const errors: Record<string, string> = {};

				e.inner.forEach((err) => {
					if (err.path && !errors[err.path]) {
						errors[err.path] = err.message;
					}
				});
				return fail(400, { errors, message: '', success: false });
			}

			return fail(500, {
				errors: { email: '', password: '' },
				message: m.form_communication_server_error(),
				success: false
			});
		}
	}
} satisfies Actions;

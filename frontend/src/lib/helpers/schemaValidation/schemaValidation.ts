import * as Yup from 'yup';
import { m } from '$lib/paraglide/messages';

export const SchemaValidation = () => {
	const username = Yup.string()
		.matches(/^[A-Z]/g, m.NavForm_validateUsernameFl())
		.matches(/[a-ząćęłńóśźżĄĘŁŃÓŚŹŻぁ-んァ-ヾ一-龯]*/g, m.NavForm_validateUsernameHKik())
		.matches(/\D/g, m.NavForm_validateUsernameNum())
		.min(3, m.NavForm_validateUsernameMin())
		.required(m.validateRequired());

	const pseudonym = Yup.string()
		.matches(/[0-9０-９]+/g, m.NavForm_validatePseudonymNum())
		.matches(/[#?!@$%^&*-＃？！＄％＆＊ー]+/g, m.NavForm_validatePseudonymSpec())
		.matches(/[a-ząćęłńóśźżĄĘŁŃÓŚŹŻぁ-んァ-ヾ一-龯]*/g, m.NavForm_validatePseudonymHKik())
		.min(5, m.NavForm_validatePseudonymMin())
		.max(15, m.NavForm_validatePseudonymMax())
		.required(m.validateRequired());

	const email = Yup.string().email(m.NavForm_validateEmail()).required(m.validateRequired());

	const password = Yup.string()
		.min(9, m.NavForm_validatePasswordNum())
		.max(72, m.NavForm_validatePasswordMax())
		.matches(/[A-Z]+/g, m.NavForm_validatePasswordOl())
		.matches(/[a-ząćęłńóśźżĄĘŁŃÓŚŹŻぁ-んァ-ヾ一-龯]*/g, m.NavForm_validatePasswordHKik())
		.matches(/[0-9]+/g, m.NavForm_validatePasswordOn())
		.matches(/[#?!@$%^&*-]+/g, m.NavForm_validatePasswordSpec())
		.required(m.validateRequired());

	const description = Yup.string()
		.min(10, m.validateMessageMin())
		.max(1500, m.validateMessageMax())
		.required(m.validateRequired());

	const name = Yup.string().required(m.validateRequired());

	const deviceName = Yup.string()
		.min(3, m.validateDeviceNameMin())
		.max(20, m.validateDeviceNameMax())
		.matches(/^[a-zA-Z0-9\s\-_ぁ-んァ-ヾ一-龯ąęłńóśźżĄĘŁŃÓŚŹŻ]+$/g, m.validateDeviceNameFormat())
		.required(m.validateRequired());

	return { username, pseudonym, email, password, description, name, deviceName };
};

const EMAIL_REG =
	/^(?:[^<>()[\]\\.,;:\s@"]+(?:\.[^<>()[\]\\.,;:\s@"]+)*|".+")@(?:\[\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}\]|(?:[a-zA-Z\-\d]+\.)+[a-zA-Z]{2,})$/;

export const validateEmail = (email: string) => EMAIL_REG.test(email);

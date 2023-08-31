import { Button, TextField, Typography } from "@mui/material";
import { makeStyles } from "@mui/styles";

const useStyles = makeStyles((theme) => ({
	formContainer: {
		display: "flex",
		flexDirection: "column",
		alignItems: "center",
		padding: theme.spacing(2),
	},
	title: {
		marginBottom: theme.spacing(2),
	},
	textField: {
		marginBottom: theme.spacing(2),
	},
	submitButton: {
		width: "150px",
	},
}));

function Form() {
	const classes = useStyles();

	function handleSubmit(event) {
		event.preventDefault();
		// 处理提交逻辑
	}

	return (
		<form class={classes.formContainer} onSubmit={handleSubmit}>
			<Typography variant="h4" class={classes.title}>
				Retalk Setup
			</Typography>
			<TextField
				label="Username"
				variant="outlined"
				class={classes.textField}
			/>
			<TextField
				label="Password"
				variant="outlined"
				type="password"
				class={classes.textField}
			/>
			<Button
				type="submit"
				variant="contained"
				color="primary"
				class={classes.submitButton}
			>
				Submit
			</Button>
		</form>
	);
}

export default Form;

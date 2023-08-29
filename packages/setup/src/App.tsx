import { makeStyles } from '@mui/styles';
import { Typography, Button, TextField } from '@mui/material';

const useStyles = makeStyles((theme) => ({
  formContainer: {
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'center',
    padding: theme.spacing(2),
  },
  title: {
    marginBottom: theme.spacing(2),
  },
  textField: {
    marginBottom: theme.spacing(2),
  },
  submitButton: {
    width: '150px',
  },
}));

function Form() {
  const classes = useStyles();

  const handleSubmit = (event) => {
    event.preventDefault();
    // 处理提交逻辑
  };

  return (
    <form className={classes.formContainer} onSubmit={handleSubmit}>
      <Typography variant="h4" className={classes.title}>
        Retalk Setup
      </Typography>
      <TextField
        label="Username"
        variant="outlined"
        className={classes.textField}
      />
      <TextField
        label="Password"
        variant="outlined"
        type="password"
        className={classes.textField}
      />
      <Button
        type="submit"
        variant="contained"
        color="primary"
        className={classes.submitButton}
      >
        Submit
      </Button>
    </form>
  );
}

export default Form;

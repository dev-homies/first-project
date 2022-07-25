import { useForm } from 'react-hook-form';
import { yupResolver } from '@hookform/resolvers/yup';
import * as yup from 'yup';
import { useMutation } from 'react-query';
import { Axios } from '../utils';
import axios from 'axios';

type LoginFormType = {
  name: string;
  password: string;
};

const schema = yup
  .object({
    name: yup.string(),
    password: yup
      .string()
      .required('Please Enter your password')
      .matches(
        /^(?=.*[A-Za-z])(?=.*\d)(?=.*[@$!%*#?&])[A-Za-z\d@$!%*#?&]{8,}$/,
        'Must Contain 8 Characters, One Uppercase, One Lowercase, One Number and one special case Character'
      ),
  })
  .required();

const LoginForm = () => {
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<LoginFormType>({
    resolver: yupResolver(schema),
  });

  const mutation = useMutation((formData: LoginFormType) => {
    return axios.post(`http://localhost:4000/v1/login`, formData);
  });

  const onUserSubmit = (data: LoginFormType) => mutation.mutate(data);

  return (
    <>
      <div className="h-screen w-screen">
        <div className="flex h-full items-center justify-center px-4">
          <form onSubmit={handleSubmit(onUserSubmit)}>
            <div className="form-control w-full">
              <label className="label">
                <span className="label-text">Name</span>
              </label>
              <input
                {...register('name')}
                placeholder="Name"
                className="input input-bordered mb-2 w-full"
              />
              {errors.name ? (
                <span className="py-2 text-sm text-red-500">{errors.name?.message}</span>
              ) : null}
            </div>
            <div className="form-control w-full">
              <label className="label">
                <span className="label-text">Password</span>
              </label>
              <input
                {...register('password')}
                type="password"
                placeholder="Password"
                className="input input-bordered mb-2 w-full"
              />
              {errors.password ? (
                <span className="py-2 text-sm text-red-500">{errors.password?.message}</span>
              ) : null}
            </div>

            <button type="submit" className="btn w-full">
              Submit
            </button>
          </form>
        </div>
      </div>
    </>
  );
};

export default LoginForm;

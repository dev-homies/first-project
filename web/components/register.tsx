import { useForm } from 'react-hook-form';
import { yupResolver } from '@hookform/resolvers/yup';
import * as yup from 'yup';
import { useMutation } from 'react-query';
import { Axios } from '../utils';

type RegisterFormType = {
  name: string;
  username: string;
  email: string;
  password: string;
};

const schema = yup
  .object({
    name: yup.string().required('Please enter your name'),
    username: yup.string().required('Please enter your username'),
    email: yup.string().email().required('Please enter your email'),
    password: yup
      .string()
      .required('Please enter your password')
      .matches(
        /^(?=.*[A-Za-z])(?=.*\d)(?=.*[@$!%*#?&])[A-Za-z\d@$!%*#?&]{8,}$/,
        'Must Contain 8 Characters, One Uppercase, One Lowercase, One Number and one special case Character'
      ),
  })
  .required();

const RegisterForm = () => {
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<RegisterFormType>({
    resolver: yupResolver(schema),
  });

  const mutation = useMutation((formData: RegisterFormType) => {
    return Axios.post('http://localhost:4000/v1/register', formData);
  });

  const onUserSubmit = (data: RegisterFormType) => mutation.mutate(data);
  return (
    <>
      <div className="h-screen w-screen">
        <div className="flex h-full items-center justify-center px-4">
          <form onSubmit={handleSubmit(onUserSubmit)} className="w-full max-w-sm">
            <div className="form-control w-full">
              <label className="label">
                <span className="label-text">Name</span>
              </label>
              <input
                {...register('name')}
                type="text"
                placeholder="Name"
                className="input input-bordered mb-2 w-full"
              />
              {errors.name ? (
                <span className="py-2 text-sm text-red-500">{errors.name?.message}</span>
              ) : null}
            </div>

            <div className="form-control w-full">
              <label className="label">
                <span className="label-text">Username</span>
              </label>
              <input
                {...register('username')}
                type="text"
                placeholder="Username"
                className="input input-bordered mb-2 w-full"
              />
              {errors.username ? (
                <span className="py-2 text-sm text-red-500">{errors.username?.message}</span>
              ) : null}
            </div>

            <div className="form-control w-full">
              <label className="label">
                <span className="label-text">Email</span>
              </label>
              <input
                {...register('email')}
                type="email"
                placeholder="Email"
                className="input input-bordered mb-2 w-full"
              />
              {errors.email ? (
                <span className="py-2 text-sm text-red-500">{errors.email?.message}</span>
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
            <button type="submit" className="btn mt-4 w-full">
              Register
            </button>
          </form>
        </div>
      </div>
    </>
  );
};

export default RegisterForm;

import { useForm } from "react-hook-form";
import { yupResolver } from "@hookform/resolvers/yup";
import * as yup from "yup";

const schema = yup
  .object({
    name: yup.string().required(),
    username: yup.string().required(),
    email: yup.string().email(),
    password: yup
      .string()
      .required("Please Enter your password")
      .matches(
        /^(?=.*[A-Za-z])(?=.*\d)(?=.*[@$!%*#?&])[A-Za-z\d@$!%*#?&]{8,}$/,
        "Must Contain 8 Characters, One Uppercase, One Lowercase, One Number and one special case Character"
      ),
  })
  .required();

const RegisterForm = () => {
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm({
    resolver: yupResolver(schema),
  });
  const callAPIFormCompleted = (data: any) => console.log(data);

  return (
    <>
      <div className="h-screen w-screen">
        <div className="flex h-full items-center justify-center px-4">
          <form onSubmit={handleSubmit(callAPIFormCompleted)}>
            <input {...register("name")} type="text" placeholder="Name" className="input input-bordered mb-2 w-full" />
            <input
              {...register("username")}
              type="text"
              placeholder="Username"
              className="input input-bordered mb-2 w-full"
            />
            <input
              {...register("email")}
              type="email"
              placeholder="Email"
              className="input input-bordered mb-2 w-full"
            />
            <input
              {...register("password")}
              type="password"
              placeholder="Password"
              className="input input-bordered mb-2 w-full"
            />
            <button type="submit" className="btn w-full">
              Submit
            </button>
          </form>
        </div>
      </div>
    </>
  );
};

export default RegisterForm;

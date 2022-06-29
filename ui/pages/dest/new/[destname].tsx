import type { NextPage } from "next";
import Head from "next/head";
import Header from "@/components/Header";
import { useRouter } from "next/router";
import DestinationFields from "@/components/DestFields";

type NewDestinationProps = {
  destname: string;
};

const NewDestination: NextPage<NewDestinationProps> = ({ destname }) => {
  const router = useRouter();
  const fields = DestinationFields[destname];

  const handleSubmit = async (event: any) => {
    event.preventDefault();
    var formData = new FormData(event.target);
    var object: { [key: string]: string } = {};
    formData.forEach(function (value, key) {
      object[key] = value.toString();
    });
    const JSONdata = JSON.stringify(object);
    console.log(JSONdata);
    const response = await fetch("/api/dests", {
      body: JSONdata,
      headers: {
        "Content-Type": "application/json",
      },
      method: "POST",
    });

    if (response.ok) {
      router.push("/");
    }
  };

  return (
    <div className="flex h-screen flex-col">
      <Head>
        <title>Observability Control Plane - UI</title>
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <Header />
      <div className="text-4xl p-8 capitalize antialiased text-gray-900">
        Add new {destname} destination
      </div>
      <div className="pl-14 max-w-md">
        <form
          className="grid grid-cols-1 gap-6"
          onSubmit={handleSubmit}
          name="newdest"
        >
          <label className="block">
            <span className="text-gray-700">Destination Name</span>
            <input
              type="text"
              id="name"
              name="name"
              className="
                    mt-1
                    block
                    w-full
                    rounded-md
                    border-gray-300
                    shadow-sm
                    focus:border-indigo-300 focus:ring focus:ring-indigo-200 focus:ring-opacity-50
                  "
              placeholder=""
              required
            />
          </label>
          {fields &&
            fields.map((f) => {
              return (
                <label className="block" key={f.id}>
                  <span className="text-gray-700">{f.displayName}</span>
                  <input
                    id={f.id}
                    name={f.name}
                    type={f.type}
                    className="
                    mt-1
                    block
                    w-full
                    rounded-md
                    border-gray-300
                    shadow-sm
                    focus:border-indigo-300 focus:ring focus:ring-indigo-200 focus:ring-opacity-50
                  "
                    placeholder=""
                    required
                  />
                </label>
              );
            })}
          {/* <label className="block">
            <span className="text-gray-700">URL</span>
            <input
              id="url"
              name="url"
              type="url"
              className="
                    mt-1
                    block
                    w-full
                    rounded-md
                    border-gray-300
                    shadow-sm
                    focus:border-indigo-300 focus:ring focus:ring-indigo-200 focus:ring-opacity-50
                  "
              placeholder=""
              required
            />
          </label>
          <label className="block">
            <span className="text-gray-700">User</span>
            <input
              type="text"
              name="user"
              id="user"
              className="
                    mt-1
                    block
                    w-full
                    rounded-md
                    border-gray-300
                    shadow-sm
                    focus:border-indigo-300 focus:ring focus:ring-indigo-200 focus:ring-opacity-50"
              required
            />
          </label>
          <label className="block">
            <span className="text-gray-700">API Key</span>
            <input
              type="password"
              name="apikey"
              id="apikey"
              className="
                    block
                    w-full
                    mt-1
                    rounded-md
                    border-gray-300
                    shadow-sm
                    focus:border-indigo-300 focus:ring focus:ring-indigo-200 focus:ring-opacity-50
                  "
              required
            />
          </label> */}
          <input name="type" id="type" hidden value={destname} readOnly />
          <button
            type="submit"
            className="mt-4 text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 mr-2 mb-2 dark:bg-blue-600 dark:hover:bg-blue-700 focus:outline-none dark:focus:ring-blue-800"
          >
            Create Destination
          </button>
        </form>
      </div>
    </div>
  );
};

export default NewDestination;

export const getServerSideProps = async ({ query }: any) => {
  return {
    props: {
      destname: query.destname,
    },
  };
};
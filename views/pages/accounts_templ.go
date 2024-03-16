// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.598
package pages

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func handleFormSubmit() templ.ComponentScript {
	return templ.ComponentScript{
		Name: `__templ_handleFormSubmit_5e67`,
		Function: `function __templ_handleFormSubmit_5e67(){document.addEventListener("DOMContentLoaded", function () {
       const signinForm = document.getElementById("signinForm");
       const loginForm = document.getElementById("loginForm");

       function handleSubmit(event) {
            const errorMessage = document.querySelector('.active-error-message');
           event.preventDefault();

           const form = event.target;
           const formData = new FormData(form);
           const jsonObject = {};

           for (const [key, value] of formData.entries()) {
               if (key.endsWith('-username') || key.endsWith('-password')) {
                   const newKey = key.split('-')[1];
                   jsonObject[newKey] = value;
               }
           }

           const finalObject = {
               user: jsonObject.username,
               password: jsonObject.password
           };

           fetch(form.action, {
               method: 'POST',
               headers: {
                   'Content-Type': 'application/json'
               },
               body: JSON.stringify(jsonObject)
           })
           .then(response => response.json().then(data => ({status: response.status, body: data})))
           .then(obj => {
               if (obj.status >= 200 && obj.status < 300) {
                   window.location.href = '/users';
               } else if (obj.body && obj.body.error) {
                console.log("Error")
                    errorMessage.classList.toggle('hidden');
                   errorMessage.childNodes[0].textContent = obj.body.error;
               } else {
                console.log("Error")
                    errorMessage.classList.toggle('hidden');
                   errorMessage.childNodes[0].textContent = "Se produjo un error en el servidor.";
               }
           })
           .catch((error) => {
                errorMessage.classList.toggle('hidden');
               errorMessage.childNodes[0].textContent = "Error en la comunicación con el servidor.";
           });
       }

       signinForm.addEventListener("submit", handleSubmit);
       loginForm.addEventListener("submit", handleSubmit);
   });
}`,
		Call:       templ.SafeScript(`__templ_handleFormSubmit_5e67`),
		CallInline: templ.SafeScriptInline(`__templ_handleFormSubmit_5e67`),
	}
}

func Accounts() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"w-full max-w-sm p-4 bg-white border border-gray-200 rounded-lg shadow sm:p-6 md:p-8 dark:bg-gray-800 dark:border-gray-700 m-auto p-10\"><h1 class=\"text-2xl font-bold text-gray-900 dark:text-white\">Create a new account</h1><form id=\"signinForm\" action=\"/api/signin\" method=\"post\"><div class=\"mb-6\"><label for=\"signin-username\" class=\"block mb-2 text-sm font-medium text-gray-900 dark:text-white\">Username</label> <input type=\"text\" id=\"signin-username\" name=\"signin-username\" class=\"bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500\" placeholder=\"Gopher123...\" required></div><div class=\"mb-6\"><label for=\"signin-password\" class=\"block mb-2 text-sm font-medium text-gray-900 dark:text-white\">Password</label> <input type=\"password\" id=\"signin-password\" name=\"signin-password\" class=\"bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500\" placeholder=\"•••••••••\" required></div><button type=\"submit\" class=\"w-full text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800\">Create account</button> <button type=\"button\" onclick=\"switchForms()\" class=\"w-full text-gray-900 bg-white border border-gray-300 focus:outline-none hover:bg-gray-100 focus:ring-4 focus:ring-gray-100 font-medium rounded-lg text-sm px-5 py-2.5 dark:bg-gray-800 dark:text-white dark:border-gray-600 dark:hover:bg-gray-700 dark:hover:border-gray-600 dark:focus:ring-gray-700 m-auto mt-4\">Already have an account? Login here</button><div class=\"signin-error-message hidden active-error-message p-2 mt-4 text-center text-sm text-red-800 rounded-lg bg-red-50 dark:bg-gray-800 dark:text-red-400\" role=\"alert\"><p class=\"font-medium\"></p></div></form><form id=\"loginForm\" action=\"/api/login\" method=\"post\" class=\"hidden\"><div class=\"mb-6\"><label for=\"login-username\" class=\"block mb-2 text-sm font-medium text-gray-900 dark:text-white\">Username</label> <input type=\"text\" id=\"login-username\" name=\"login-username\" class=\"bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500\" placeholder=\"Gopher123...\" required></div><div class=\"mb-6\"><label for=\"login-password\" class=\"block mb-2 text-sm font-medium text-gray-900 dark:text-white\">Password</label> <input type=\"password\" id=\"login-password\" name=\"login-password\" class=\"bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500\" placeholder=\"•••••••••\" required></div><button type=\"submit\" class=\"w-full text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800\">Login into your account</button> <button type=\"button\" onclick=\"switchForms()\" class=\"w-full text-gray-900 bg-white border border-gray-300 focus:outline-none hover:bg-gray-100 focus:ring-4 focus:ring-gray-100 font-medium rounded-lg text-sm px-5 py-2.5 dark:bg-gray-800 dark:text-white dark:border-gray-600 dark:hover:bg-gray-700 dark:hover:border-gray-600 dark:focus:ring-gray-700 m-auto mt-4\">Don't have an account? Sign in here</button><div class=\"login-error-message hidden p-2 mt-4 text-center text-sm text-red-800 rounded-lg bg-red-50 dark:bg-gray-800 dark:text-red-400\" role=\"alert\"><p class=\"font-medium\"></p></div></form></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = handleFormSubmit().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<script>\n    let signInForm = document.querySelector(\"#signinForm\");\n    let loginForm = document.querySelector(\"#loginForm\");\n\n    let signinErrorMessage = document.querySelector(\".signin-error-message\");\n    let loginErrorMessage = document.querySelector(\".login-error-message\");\n\n    function switchForms() {\n        signInForm.classList.toggle('hidden');\n        loginForm.classList.toggle('hidden');\n\n        signinErrorMessage.classList.toggle('active-error-message');\n        loginErrorMessage.classList.toggle('active-error-message');\n    }\n    </script>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

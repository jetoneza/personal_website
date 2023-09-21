<div class="container login mx-auto p-6 flex flex-col justify-center items-center h-screen">
	<div class="form-wrapper w-80 p-6 rounded-lg drop-shadow-lg bg-white">
		<div class="title font-bold text-lg">Login to Dashboard</div>
		<form class="flex flex-col space-y-4 mt-4">
			<div class="input-wrapper">
				<label for="email" class="block mb-2 text-sm font-medium dark:text-white"
					>Email address</label
				>
				<input
					type="email"
					id="email"
					class="
            border border-gray-300 rounded-lg text-sm block w-full p-2.5
            focus:ring-2 focus:outline-none focus:ring-blue-300
          "
					placeholder="admin@company.com"
					required
				/>
			</div>
			<div class="input-wrapper">
				<label for="password" class="block mb-2 text-sm font-medium dark:text-white">
					Password
				</label>
				<input
					type="password"
					id="password"
					class="
            border border-gray-300 rounded-lg text-sm block w-full p-2.5
            focus:ring-2 focus:outline-none focus:ring-blue-300
          "
					placeholder="**********"
					required
				/>
			</div>
			<button
				type="submit"
				class="
          text-white bg-cyan-600 hover:bg-cyan-700 focus:ring-4
          focus:outline-none focus:ring-blue-300 font-medium rounded-lg
          text-sm w-full sm:w-auto px-5 py-2.5 text-center
        "
			>
				Submit
			</button>
		</form>
	</div>
</div>

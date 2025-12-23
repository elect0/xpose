<script>
	import EduImage from '$lib/assets/edu.png';
</script>

<div class="grid-bg flex max-h-screen justify-center bg-[#CFFB80] pt-40 text-center">
	<div class="flex flex-col items-center space-y-5 px-4">
		<div
			class="max-w-2xl rounded-full border-2 border-black bg-[#FFFBEF] px-4 py-2 text-sm font-semibold text-black shadow-[2px_2px_0_0_#000]"
		>
The Ultimate Friendship Ruining App
		</div>

		<h1 class="text-7xl font-bold">
			Vote, Roast, Expose, Laugh. All in <br> One App.
		</h1>
		<p class="max-w-4xl text-lg">
			Forget safe conversations. Every 24 hours, a new question drops. Accuse your friends, lock in
			your votes, and check the receipts to see who the group really thinks is the problem.
		</p>
		<button
			class="transform cursor-pointer rounded-full border-2 border-black bg-[#FFFBEF] px-6 py-3 text-lg font-semibold text-black shadow-[2px_2px_0_0_#000] transition-all hover:translate-x-1 hover:translate-y-1 hover:shadow-none active:translate-x-1 active:translate-y-1 active:shadow-none"
		>
			Enter the lobby.
		</button>
		<div class="mt-4 ml-8 h-[px] overflow-hidden">
			<enhanced:img src={EduImage ? EduImage : ''} alt="Phone image" class="w-64 md:w-80 lg:w-96" />
		</div>
	</div>
</div>

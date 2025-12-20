<script>
	import { browser } from '$app/environment';
	import Scene from '$lib/components/scene.svelte';
</script>

<div class="relative flex min-h-screen items-center justify-center overflow-hidden bg-[#FAF5EC]">
	<div class="relative z-10 container mx-auto px-4 pt-28 lg:px-6">
		<h1 class="mb-6 text-center text-8xl font-bold text-[#1468C1]">Time to call 'em out.'</h1>
		<div class="svg-grid grid grid-cols-2 gap-8 p-12">
			<div class="block space-y-4">
				<div
					class="rounded-md border-3 border-black bg-[#FAF5EC] px-5 py-3 text-sm font-semibold text-black shadow-[4px_4px_0_0_#F25C69] md:text-lg"
				>
					[Player 1]
				</div>

				<div
					class="rounded-md border-3 border-black bg-[#FAF5EC] px-5 py-3 text-sm font-semibold text-black shadow-[4px_4px_0_0_#F25C69] md:text-lg"
				>
					[Player 1]
				</div>

				<div
					class="rounded-md border-3 border-black bg-[#FAF5EC] px-5 py-3 text-sm font-semibold text-black shadow-[4px_4px_0_0_#F25C69] md:text-lg"
				>
					[Player 1]
				</div>

				<button
					class="w-full transform cursor-pointer rounded-md border-2 border-black bg-[#F25C69] px-5 py-3 text-lg font-semibold text-black shadow-[4px_4px_0_0_#000] transition-all hover:translate-x-1 hover:translate-y-1 hover:shadow-none active:translate-x-1 active:translate-y-1 active:shadow-none"
				>
					Let's Play!
				</button>
			</div>
			<img src="/controller.png" alt="" />
		</div>
	</div>
</div>

<style>
	.hero-3d {
		width: 100%;
		height: 300px; /* Ensure the parent has height */
		border-radius: 1rem;
		overflow: hidden;
	}
</style>

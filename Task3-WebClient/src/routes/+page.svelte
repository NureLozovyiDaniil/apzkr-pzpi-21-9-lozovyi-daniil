<script>
	import { SlidersHorizontal, BadgeCheck, ScanEye, Handshake } from 'lucide-svelte';
	import { Button } from '$lib/components/ui/button';
	import * as Card from '$lib/components/ui/card';
	import { goto } from '$app/navigation';
</script>

<div class="bg-primary min-h-screen">
	<div class="container mx-auto px-4 py-12">
		<div class="flex flex-col items-center justify-between lg:flex-row">
			<div class="mb-8 flex max-w-4xl flex-col justify-between text-white lg:mb-0">
				<div>
					<div class="mb-2 text-center text-lg">Відслідковуйте рух гуманітрної допомоги</div>
					<div class="mb-8 text-center text-5xl font-bold">Прозоро та Ефективно</div>
				</div>
				<div class="text-center">
					<p class="mb-6 text-center">
						Інноваційні рішення для управління та відстеження гуманітарних вантажів Отримайте
						потужні інструменти для організації, відстеження та розподілу гуманітарної допомоги.
						Забезпечте прозорість процесу, мінімізуйте ризики втрат і максимізуйте ефективність
						вашої гуманітарної діяльності.
					</p>
					<Button variant="secondary" class="px-4 py-2 text-lg">Заявки на допомогу</Button>
				</div>
			</div>
			<div class="w-full max-w-md rounded-lg bg-white shadow-md">
				<div class="p-6">
					<form>
						<div class="space-y-4">
							<input class="w-full rounded border px-3 py-2" placeholder="Повне ім'я" />
							<input
								type="email"
								class="w-full rounded border px-3 py-2"
								placeholder="Email адреса"
							/>
							<select class="w-full rounded border px-3 py-2">
								<option value="">Я потребую допомогу</option>
								<option value="">Я волонтер</option>
								<option value="">Я представляю організацію</option>
								<!-- Add more options here -->
							</select>
							<Button
								class="w-full rounded bg-blue-600 px-4 py-2 text-white"
								on:click={() => goto('/register')}>Зареєструватись</Button
							>
						</div>
					</form>
				</div>
			</div>
		</div>
	</div>
	<div class="bg-white py-16">
		<div class="container mx-auto px-4">
			<h2 class="mb-4 text-center text-3xl font-bold">Для чого AidTracker?</h2>
			<p class="mx-auto mb-12 max-w-2xl text-center">
				Інноваційні рішення для управління та відстеження гуманітарних вантажів
			</p>
			<div class="grid grid-cols-1 gap-8 md:grid-cols-3">
				<Card.Root>
					<Card.Header>
						<Card.Title class="flex items-center text-2xl">
							<SlidersHorizontal class="text-primary " />&nbsp;Управління
						</Card.Title>
						<Card.Description>Ефективна організація гуманітарної допомоги</Card.Description>
					</Card.Header>
					<Card.Content>
						<ul class="text-md space-y-3">
							<li class="flex items-start">
								<BadgeCheck class="mr-2 mt-0.5 h-5 w-5 flex-shrink-0 text-green-500" />&nbsp;<span
									class="flex-grow text-justify">Створюйте та переглядайте заявки на допомогу</span
								>
							</li>
							<li class="flex items-start">
								<BadgeCheck class="mr-2 mt-0.5 h-5 w-5 flex-shrink-0 text-green-500" />&nbsp;
								<span class="flex-grow text-justify"
									>Контролюйте весь процес від збору до доставки</span
								>
							</li>
							<li class="flex items-start">
								<BadgeCheck class="mr-2 mt-0.5 h-5 w-5 flex-shrink-0 text-green-500" />&nbsp;
								<span class="flex-grow text-justify"
									>Призначайте відповідальних осіб для своїх вантажів</span
								>
							</li>
						</ul>
					</Card.Content>
				</Card.Root>
				<Card.Root>
					<Card.Header>
						<Card.Title class="flex items-center text-2xl">
							<ScanEye class="text-primary " />&nbsp;Відстеження
						</Card.Title>
						<Card.Description>Повна прозорість руху вантажів</Card.Description>
					</Card.Header>
					<Card.Content>
						<ul class="text-md space-y-3">
							<li class="flex items-start">
								<BadgeCheck class="mr-2 mt-0.5 h-5 w-5 flex-shrink-0 text-green-500" />&nbsp;<span
									class="flex-grow text-justify"
									>Відстежуйте місце розташування і статус вантажів у реальному часі</span
								>
							</li>
							<li class="flex items-start">
								<BadgeCheck class="mr-2 mt-0.5 h-5 w-5 flex-shrink-0 text-green-500" />&nbsp;
								<span class="flex-grow text-justify"
									>Отримуйте інформацію про ключові етапи доставки</span
								>
							</li>
							<li class="flex items-start">
								<BadgeCheck class="mr-2 mt-0.5 h-5 w-5 flex-shrink-0 text-green-500" />&nbsp;
								<span class="flex-grow text-justify"
									>Вимагайте фото або RFID підвтердження на кожному етапі</span
								>
							</li>
						</ul>
					</Card.Content>
				</Card.Root>
				<Card.Root>
					<Card.Header>
						<Card.Title class="flex items-center text-2xl">
							<Handshake class="text-primary " />&nbsp;Співпраця
						</Card.Title>
						<Card.Description>Об'єднання зусиль для максимального ефекту</Card.Description>
					</Card.Header>
					<Card.Content>
						<ul class="text-md space-y-3">
							<li class="flex items-start">
								<BadgeCheck class="mr-2 mt-0.5 h-5 w-5 flex-shrink-0 text-green-500" />&nbsp;<span
									class="flex-grow text-justify"
									>Приймайте волонтерів до організацій та координуйте їх дії</span
								>
							</li>
							<li class="flex items-start">
								<BadgeCheck class="mr-2 mt-0.5 h-5 w-5 flex-shrink-0 text-green-500" />&nbsp;
								<span class="flex-grow text-justify"
									>Поєднуйте нужденних із волонтерами чи організаціями, готовими допомогти</span
								>
							</li>
							<li class="flex items-start">
								<BadgeCheck class="mr-2 mt-0.5 h-5 w-5 flex-shrink-0 text-green-500" />&nbsp;
								<span class="flex-grow text-justify"
									>Обмінюйтеся досвідом і ресурсами в єдиному просторі</span
								>
							</li>
						</ul>
					</Card.Content>
				</Card.Root>
			</div>
		</div>
	</div>
</div>

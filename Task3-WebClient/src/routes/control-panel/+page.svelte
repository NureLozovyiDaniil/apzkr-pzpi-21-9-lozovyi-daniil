<script lang="ts">
	import { onMount } from 'svelte';
	import { Button } from '$lib/components/ui/button';
	import { Input } from '$lib/components/ui/input';
	import { Textarea } from '$lib/components/ui/textarea';
	import { Label } from '$lib/components/ui/label';
	import { User, Shield, Package, Users, UserPlus, Building } from 'lucide-svelte';

	let userRole = getCookie('userRole') || 'guest';
	let activeSection: 'cargo' | 'members' | 'volunteers' | 'users' | 'organizations';
	if (userRole === '3') {
		activeSection = 'cargo';
	} else if (userRole === '4') {
		activeSection = 'users';
	}

	function getCookie(name: string) {
		const value = `; ${document.cookie}`;
		const parts = value.split(`; ${name}=`);
		if (parts.length === 2) return parts.pop()?.split(';').shift();
	}

	// Placeholder functions for new sections
	function handleCargo() {
		console.log('Cargo section clicked');
	}

	function handleMembers() {
		console.log('Members section clicked');
	}

	function handleVolunteers() {
		console.log('Volunteers section clicked');
	}

	function handleUsers() {
		console.log('Users section clicked');
	}

	function handleOrganizations() {
		console.log('Organizations section clicked');
	}
</script>

<div class="container mx-auto p-4">
	<h1 class="mb-4 text-2xl font-bold">Control panel</h1>
	<p class="mb-6 text-gray-600">Main service functions</p>

	<div class="flex">
		<!-- Left sidebar -->
		<div class="w-1/4 pr-4">
			{#if userRole === '3'}
				<Button
					variant={activeSection === 'cargo' ? 'default' : 'ghost'}
					class="mb-2 w-full justify-start"
					on:click={() => (activeSection = 'cargo')}
				>
					<Package class="mr-2 h-4 w-4" />
					Cargo
				</Button>
				<Button
					variant={activeSection === 'members' ? 'default' : 'ghost'}
					class="mb-2 w-full justify-start"
					on:click={() => (activeSection = 'members')}
				>
					<Users class="mr-2 h-4 w-4" />
					Members
				</Button>
				<Button
					variant={activeSection === 'volunteers' ? 'default' : 'ghost'}
					class="mb-2 w-full justify-start"
					on:click={() => (activeSection = 'volunteers')}
				>
					<UserPlus class="mr-2 h-4 w-4" />
					Volunteers
				</Button>
			{:else if userRole === '4'}
				<Button
					variant={activeSection === 'users' ? 'default' : 'ghost'}
					class="mb-2 w-full justify-start"
					on:click={() => (activeSection = 'users')}
				>
					<Users class="mr-2 h-4 w-4" />
					Users
				</Button>
				<Button
					variant={activeSection === 'organizations' ? 'default' : 'ghost'}
					class="mb-2 w-full justify-start"
					on:click={() => (activeSection = 'organizations')}
				>
					<Building class="mr-2 h-4 w-4" />
					Organizations
				</Button>
			{/if}
		</div>

		<!-- Main content -->
		<div class="w-3/4">
			{#if activeSection === 'cargo'}
				<h2 class="mb-4 text-xl font-semibold">Cargo Management</h2>
				<Button on:click={handleCargo}>Manage Cargo</Button>
			{:else if activeSection === 'members'}
				<h2 class="mb-4 text-xl font-semibold">Member Management</h2>
				<Button on:click={handleMembers}>Manage Members</Button>
			{:else if activeSection === 'volunteers'}
				<h2 class="mb-4 text-xl font-semibold">Volunteer Management</h2>
				<Button on:click={handleVolunteers}>Manage Volunteers</Button>
			{:else if activeSection === 'users'}
				<h2 class="mb-4 text-xl font-semibold">User Management</h2>
				<Button on:click={handleUsers}>Manage Users</Button>
			{:else if activeSection === 'organizations'}
				<h2 class="mb-4 text-xl font-semibold">Organization Management</h2>
				<Button on:click={handleOrganizations}>Manage Organizations</Button>
			{/if}
		</div>
	</div>
</div>

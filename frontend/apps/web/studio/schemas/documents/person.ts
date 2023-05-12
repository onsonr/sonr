import {RiUserSmileLine} from 'react-icons/ri';
import {defineType, defineField} from 'sanity';

const person = defineType({
	name: 'person',
	type: 'document',
	title: 'Persons',
	icon: RiUserSmileLine,
	fields: [
		defineField({
			name: 'name',
			title: 'Name',
			type: 'string',
			validation: (Rule) => Rule.required()
		}),
		defineField({
			name: 'title',
			title: 'Job title',
			type: 'string',
			validation: (Rule) => Rule.required()
		}),
		defineField({
			name: 'phone',
			title: 'Phone',
			type: 'string',
			validation: (Rule) => Rule.required()
		}),
		defineField({
			name: 'email',
			title: 'email',
			type: 'email',
			validation: (Rule) => Rule.required()
		}),
		defineField({
			name: 'image',
			title: 'Image',
			type: 'image',
			options: {
				hotspot: true
			},
			validation: (Rule) => Rule.required()
		})
	],
	preview: {
		select: {
			title: 'name',
			media: 'image.asset'
		}
	}
});

export default person;

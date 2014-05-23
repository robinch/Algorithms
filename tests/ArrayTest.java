import java.util.Arrays;

class ArrayTest{
	public ArrayTest(){

		int[] array = {1,2,3};
		System.out.println("Array Test");
		System.out.println(Arrays.toString(array));
		changeArray(array);
		System.out.println(Arrays.toString(array));
		System.out.println("Value Test");
		int value = 1;
		System.out.println(value);
		changeValue(value);
		System.out.println(value);


	}

	public void changeArray(int[] a){
		a[2] = 75;
	}

	public void changeValue(int v){
		v = 75;
	}

	public static void main(String[] args){
		System.out.println("Test");
		ArrayTest at = new ArrayTest();
		
	}
}
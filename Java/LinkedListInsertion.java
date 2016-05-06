public class LinkedListInsertion<T> {
	
	Node<T> head;
	
	public static class Node<T>{
		T data;
		Node<T> next;
		
		Node(T data){
			this.data = data;
			this.next = null;
		}
	}
	
	
	
	//Insert node at tail
	public void append(T data){
		
		Node<T> newTailNode = new Node<T>(data);
		
		if(null == head){
			head = newTailNode;
			return;
		}
		
		Node<T> last = head;
		while(null != last.next){
			last = last.next;
		}
		
		last.next = newTailNode;
			
	}
	
	
	//Insert node as head 
		public void push(T data){
			Node<T> newHeadNode = new Node<T>(data);
			newHeadNode.next = head;
			head = newHeadNode;
		}
	
	
	void insertAfter(T data,Node<T> previous){
		if(null == previous){
			return;
		}
		
		Node<T> newNode = new Node<T>(data);
		newNode.next = previous.next;
		previous.next = newNode;
		
	}
	
	
	 void printList(){
		Node<T> node = head;
		while(null != node){
			System.out.println(node.data);
			node = node.next;
		}
	}
	
	
	public static void main(String[] args) {
		LinkedListInsertion<Integer> list = new LinkedListInsertion<Integer>();
		
		list.head = new Node<Integer>(10); //Create head node
		
		Node<Integer> node2 = new Node<Integer>(20);
		list.head.next = node2; //Set node2 as next of head node
		
		Node<Integer> node3 = new Node<Integer>(30);
		node2.next = node3; //Set node 3 as next of node 2
		
		list.push(5);
		
		list.append(40);
		
		list.insertAfter(35, node3);
		
		
		list.printList();
		
	}
	
}

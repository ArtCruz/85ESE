import React from 'react';
import Table from 'react-bootstrap/Table';
import Modal from 'react-bootstrap/Modal';
import Button from 'react-bootstrap/Button';
import Form from 'react-bootstrap/Form';
import axios from 'axios';

class ProductList extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            products: [],
            showAddProductDialog: false,
            isEditing: false,
            newProduct: {
                id: '',
                name: '',
                price: '',
                sku: '',
            },
        };

        this.readData = this.readData.bind(this);
        this.showAddProductDialog = this.showAddProductDialog.bind(this);
        this.closeAddProductDialog = this.closeAddProductDialog.bind(this);
        this.handleInputChange = this.handleInputChange.bind(this);
        this.handleSaveProduct = this.handleSaveProduct.bind(this);
    }

    componentDidMount() {
        this.readData();
    }

    readData() {
        axios.get(window.global.gateway_location + '/products')
            .then((response) => {
                console.log(response.data);
                this.setState({ products: response.data });
            })
            .catch((error) => {
                console.log(error);
            });
    }

    getProducts() {
        let table = [];

        for (let i = 0; i < this.state.products.length; i++) {
            table.push(
                <tr key={i}>
                    <td>{this.state.products[i].id}</td>
                    <td>{this.state.products[i].name}</td>
                    <td>{this.state.products[i].description}</td>
                    <td>{this.state.products[i].price}</td>
                    <td>{this.state.products[i].sku}</td>
                    <td>
                        <button
                            className="btn btn-secondary"
                            onClick={() => this.editProduct(this.state.products[i])}
                        >
                            Edit
                        </button>
                        <button className="btn btn-danger" onClick={() => this.deletarProduto(this.state.products[i].id)}>Delete</button>
                    </td>
                </tr>
            );
        }

        return table;
    }

    deletarProduto(productId) {
        if (window.confirm("Tem certeza que deseja deletar este produto?")) {
            axios.delete(window.global.gateway_location + `/products/${productId}`)
                .then((response) => {
                    console.log("Produto deletado com sucesso:", response.data);
                    this.readData(); // Atualiza a lista de produtos
                })
                .catch((error) => {
                    console.error("Erro ao deletar produto:", error);
                });
        }
    }

    showAddProductDialog() {
        this.setState({ showAddProductDialog: true });
    }

    closeAddProductDialog() {
        this.setState({
            showAddProductDialog: false,
            isEditing: false,
            newProduct: {
                id: '',
                name: '',
                price: '',
                sku: '',
            },
        });
    }

    handleInputChange(event) {
        const { name, value } = event.target;
        this.setState((prevState) => ({
            newProduct: {
                ...prevState.newProduct,
                [name]: value,
            },
        }));
    }

    handleSaveProduct() {
        const { newProduct, isEditing } = this.state;

        if (isEditing) {
            axios.put(window.global.gateway_location + `/products/${newProduct.id}`, newProduct)
                .then((response) => {
                    const data = response.data;

                    console.log('Resposta do backend:', data);

                    if (data.messages && Array.isArray(data.messages)) {
                        const errorMessages = data.messages;

                        if (errorMessages.some(msg => msg.includes("Product.SKU") && msg.includes("validation for 'SKU' failed"))) {
                            alert("Formatação errada de SKU");
                        } else {
                            alert("Erro desconhecido:\n" + errorMessages.join('\n'));
                        }

                        console.error('Erro ao salvar produto:', errorMessages);
                        return; // Impede a continuação
                    }

                    // Se chegou aqui, a resposta é considerada válida
                    console.log('Produto adicionado com sucesso:', data);
                    this.closeAddProductDialog();
                    this.readData();
                    this.state.showAddProductDialog = false;
                })
        } else {
            const { id, price, ...rest } = this.state.newProduct;
            const productToSend = {
                ...rest,
                price: parseFloat(price)
            };

            axios.post(window.global.gateway_location + '/products', productToSend)
                .then((response) => {
                    console.log('Produto adicionado com sucesso:', response.data);
                    this.closeAddProductDialog();
                    this.readData();
                    this.state.showAddProductDialog = false;
                })
                .catch((error) => {
                    if (error.response && error.response.data && error.response.data.Messages) {
                        alert(error.response.data.Messages.join('\n'));
                    } else if (error.response && error.response.data && error.response.data.message) {
                        alert(error.response.data.message);
                    } else {
                        alert('Formatação errada do SKU');
                    }
                    console.error('Erro ao adicionar produto:', error);
                });
        }
    }

    editProduct(product) {
        this.setState({
            newProduct: { ...product },
            showAddProductDialog: true,
            isEditing: true,
        });
    }

    render() {
        return (
            <div>
                <h1 style={{ marginBottom: '40px' }}>Menu</h1>
                <div style={{ display: 'flex', padding: '5rem', width: '80%', backgroundColor: 'red', margin: '0 auto' }}>
                    <Table>
                        <thead>
                            <tr>
                                <th>ID</th>
                                <th>Name</th>
                                <th>Description</th>
                                <th>Price</th>
                                <th>SKU</th>
                                <th>Options</th>
                            </tr>
                        </thead>
                        <tbody>
                            {this.getProducts()}
                        </tbody>
                        <tfoot>
                            <tr>
                                <td colSpan="6">
                                    <button
                                        className="btn btn-primary"
                                        onClick={this.showAddProductDialog}
                                    >
                                        Add Product
                                    </button>
                                </td>
                            </tr>
                        </tfoot>
                    </Table>

                    <Modal
                        show={this.state.showAddProductDialog}
                        onHide={this.closeAddProductDialog}
                    >
                        <Modal.Header closeButton>
                            <Modal.Title>
                                {this.state.isEditing ? 'Edit Product' : 'Add Product'}
                            </Modal.Title>
                        </Modal.Header>
                        <Modal.Body>
                            <Form>
                                <Form.Group controlId="productName">
                                    <Form.Label>Name</Form.Label>
                                    <Form.Control
                                        type="text"
                                        name="name"
                                        placeholder="Enter product name"
                                        value={this.state.newProduct.name}
                                        onChange={this.handleInputChange}
                                    />
                                </Form.Group>
                                <Form.Group controlId="productDescription">
                                    <Form.Label>Description</Form.Label>
                                    <Form.Control
                                        type="text"
                                        name="description"
                                        placeholder="Enter product description"
                                        value={this.state.newProduct.description}
                                        onChange={this.handleInputChange}
                                    />
                                </Form.Group>
                                <Form.Group controlId="productPrice">
                                    <Form.Label>Price</Form.Label>
                                    <Form.Control
                                        type="number"
                                        name="price"
                                        placeholder="Enter product price"
                                        value={this.state.newProduct.price}
                                        onChange={this.handleInputChange}
                                    />
                                </Form.Group>
                                <Form.Group controlId="productSKU">
                                    <Form.Label>SKU</Form.Label>
                                    <Form.Control
                                        type="text"
                                        name="sku"
                                        placeholder="Enter product SKU"
                                        value={this.state.newProduct.sku}
                                        onChange={this.handleInputChange}
                                    />
                                </Form.Group>
                            </Form>
                        </Modal.Body>
                        <Modal.Footer>
                            <Button variant="secondary" onClick={this.closeAddProductDialog}>
                                Close
                            </Button>
                            <Button variant="primary" onClick={this.handleSaveProduct}>
                                Save Changes
                            </Button>
                        </Modal.Footer>
                    </Modal>
                </div>
            </div>
        );
    }
}

export default ProductList;
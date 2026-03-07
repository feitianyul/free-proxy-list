**中文** | [English (README_EN.md)](README_EN.md)

---

<p align="center">
  <img src="https://img.shields.io/badge/每1小时更新-通过-success">  
  <br>
  <img src="https://img.shields.io/website/https/getfreeproxy.com.svg">
  <img src="https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/total.svg">
  <img src="https://img.shields.io/github/last-commit/feitianyul/free-proxy-list.svg">
  <img src="https://img.shields.io/github/license/feitianyul/free-proxy-list.svg">
  
  <br>
  <br>
  <a href="https://getfreeproxy.com/lists/" title="可用代理列表">可用代理列表</a> | <a href="https://getfreeproxy.com/tools/proxy-checker" title="在线代理检测">免费代理检测</a> | <a href="https://getfreeproxy.com/tools/proxy-protocol-parser" title="代理协议解析">通用代理协议解析</a> | <a href="https://developer.getfreeproxy.com/" title="代理 API">免费代理 API</a>
  <br>
</p>

# 🌎 GetFreeProxy (GFP)：免费代理列表

**GetFreeProxy (GFP)** 是一个开源项目，自动从互联网聚合并校验免费代理，旨在为开发者、研究人员及需要代理服务的用户提供新鲜、可靠、可用的公共代理列表。

列表按小时更新，确保您始终能获取到最新的可用代理。

---

## 📖 项目说明

本项目为开源免费代理聚合与校验工具，从互联网公开源拉取代理并**仅保留 HTTP、HTTPS** 两种类型，经校验后生成列表，供开发者、研究人员等使用。

### 本仓库特点

- **仅保留两种代理**：HTTP、HTTPS，不收录 SOCKS、VMess、Trojan、VLESS、SS/SSR、Hysteria 等其它协议。
- **校验规则**：五域名中**任意 3 个**在 2 秒内成功（HTTP 200）即视为该协议通过。对每条代理访问以下五个地址验证（优先 HEAD，不支持则回退 GET）：
  - `https://www.eastmoney.com/`
  - `https://www.sse.com.cn/`
  - `https://finance.sina.com.cn/`（新浪财经）
  - `https://web.ifzq.gtimg.cn/`
  - `https://proxy.finance.qq.com/`
  每个代理分别以 **HTTP 代理** 和 **HTTPS 代理** 各测一次；**协议** 写入 meta：只通 HTTP→`http`，只通 HTTPS→`https`，两个都通→`http/s`。去重按 **协议+IP+端口**。校验时**多代理并发**、**单代理内五域名并行**。列表直显：表格列为「代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议」。
- **更新频率**：列表按小时更新，保证可用代理的时效性。
- **并发参数**：校验 worker 数可通过 `-check-workers`（如 `-check-workers=4000`）或环境变量 `GFP_CHECK_WORKERS` 设置，默认 4000，最大 4000。遇目标站限流可适当调低。

### 工作流程

1. **拉取**：从 `sources/` 目录下配置的源（仅处理 `http.txt`、`https.txt`）拉取原始代理数据，支持动态 URL 及 Base64 等格式。
2. **解析与规范化**：将原始数据解析为标准代理格式（协议、IP、端口、认证等）。
3. **校验**：对 HTTP/HTTPS 代理通过上述验证与 2 秒超时规则进行筛选。
4. **去重与存储**：通过校验的代理去重后写入内存。
5. **生成列表**：按协议生成 `list/` 目录下的 `http.txt`、`https.txt`，并更新统计与 README 中的下载表格。

自动化由 GitHub Actions 执行：**全量流程**（抓取→解析→验证→生成列表）**每 6 小时**运行一次；**轻量复测**（对已有列表做连通性复测、剔除失效代理）**每 1 小时**运行一次。全量任务最长运行 12 小时，超时才会取消。下表「最后更新」时间为 UTC 及 UTC+8。

### 支持的代理格式示例

| 类型 | 格式 | 示例 |
| :--- | :--- | :--- |
| **HTTP/HTTPS** | `http://ip:port` | `http://1.2.3.4:8080` |
| | `https://ip:port` | `https://1.2.3.4:8080` |
| | `http://user:pass@ip:port` | `http://user:pass@1.2.3.4:8080` |

---

## 🔗 直接下载链接

点击下方表格中您需要的协议类型即可获取最新列表，链接始终指向最近更新的代理文件。上述三个文件（http.txt、https.txt、passed.txt）同时发布至 [GitHub Releases (tag: lists)](https://github.com/feitianyul/free-proxy-list/releases/tag/lists)，每次运行覆盖同版本附件，可固定使用该 Release 的附件 URL。

<!-- BEGIN PROXY LIST -->

最后更新：2026-03-07 14:53:34 UTC（2026-03-07 22:53:34 UTC+8）

**代理总数：43**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 42 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 43 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 1.231.81.166:3128 | ✓ 1378ms | ✓ 1473ms | ✓ 951ms | ✓ 1026ms | ✓ 743ms | http |
| 205.209.118.30:3138 | 否 | ✓ 1175ms | ✓ 1209ms | ✓ 1294ms | ✓ 1218ms | http |
| 159.223.42.219:3128 | 否 | 否 | ✓ 888ms | ✓ 1185ms | ✓ 905ms | http |
| 35.225.22.61:80 | ✓ 794ms | 否 | ✓ 486ms | 否 | ✓ 894ms | http |
| 120.92.212.16:7890 | ✓ 981ms | 否 | 否 | ✓ 1248ms | ✓ 982ms | http |
| 61.72.221.94:3128 | ✓ 856ms | 否 | ✓ 1372ms | ✓ 1202ms | ✓ 1062ms | http |
| 165.227.5.10:8888 | ✓ 1363ms | 否 | ✓ 121ms | ✓ 974ms | ✓ 1740ms | http |
| 120.92.212.16:8890 | 否 | ✓ 1224ms | ✓ 1038ms | ✓ 1496ms | 否 | http |
| 178.236.245.59:3128 | ✓ 861ms | ✓ 1870ms | ✓ 1671ms | 否 | 否 | http |
| 103.84.95.54:7890 | 否 | 否 | ✓ 1221ms | ✓ 1362ms | ✓ 1914ms | http |
| 46.183.25.8:443 | 否 | 否 | ✓ 1174ms | ✓ 1440ms | ✓ 1175ms | http |
| 8.219.97.248:80 | 否 | 否 | ✓ 1134ms | ✓ 1529ms | ✓ 1232ms | http |
| 147.45.251.242:8888 | ✓ 1188ms | ✓ 1896ms | ✓ 1694ms | 否 | 否 | http |
| 81.70.169.194:80 | 否 | ✓ 1911ms | ✓ 1996ms | ✓ 1199ms | ✓ 1289ms | http |
| 14.225.217.30:7890 | ✓ 935ms | 否 | ✓ 866ms | 否 | ✓ 1657ms | http |
| 101.43.255.96:80 | ✓ 1004ms | ✓ 1235ms | 否 | ✓ 1331ms | ✓ 1316ms | http |
| 106.14.203.63:3333 | ✓ 819ms | ✓ 1878ms | ✓ 926ms | ✓ 1920ms | ✓ 871ms | http |
| 178.236.245.17:3128 | ✓ 813ms | 否 | ✓ 1572ms | 否 | ✓ 1510ms | http |
| 91.193.240.157:9877 | ✓ 1151ms | 否 | ✓ 1060ms | 否 | ✓ 1625ms | http |
| 39.104.201.40:7890 | 否 | 否 | ✓ 1625ms | ✓ 1240ms | ✓ 957ms | http |
| 168.235.110.63:3128 | ✓ 340ms | ✓ 1288ms | ✓ 1887ms | ✓ 1207ms | ✓ 973ms | http |
| 14.56.107.244:3128 | ✓ 1512ms | 否 | 否 | ✓ 1627ms | ✓ 920ms | http |
| 121.128.121.54:3128 | 否 | 否 | ✓ 1674ms | ✓ 1100ms | ✓ 1864ms | http |
| 45.136.198.40:3128 | ✓ 772ms | 否 | ✓ 897ms | ✓ 1655ms | ✓ 1279ms | http |
| 121.40.231.103:7890 | ✓ 911ms | ✓ 1065ms | 否 | 否 | ✓ 1913ms | http |
| 125.128.12.144:3128 | ✓ 1757ms | ✓ 1383ms | ✓ 984ms | ✓ 1035ms | ✓ 848ms | http |
| 61.72.110.54:3128 | 否 | 否 | ✓ 1135ms | ✓ 1121ms | ✓ 978ms | http |
| 167.172.69.123:8080 | ✓ 743ms | 否 | ✓ 1239ms | ✓ 1103ms | ✓ 853ms | http |
| 167.172.69.123:80 | ✓ 737ms | 否 | ✓ 1172ms | ✓ 1071ms | ✓ 872ms | http |
| 61.72.221.194:3128 | ✓ 1799ms | 否 | ✓ 1525ms | ✓ 1080ms | 否 | http |
| 61.72.221.234:3128 | 否 | ✓ 1865ms | ✓ 923ms | 否 | ✓ 1520ms | http |
| 14.56.177.44:3128 | 否 | 否 | ✓ 1327ms | ✓ 1792ms | ✓ 1082ms | http |
| 103.139.138.194:3128 | 否 | 否 | ✓ 1433ms | ✓ 1778ms | ✓ 1158ms | http |
| 91.233.223.147:3128 | ✓ 1371ms | 否 | ✓ 1262ms | 否 | ✓ 1587ms | http |
| 173.212.246.157:3128 | 否 | 否 | ✓ 1486ms | ✓ 1999ms | ✓ 1553ms | http |
| 115.231.181.40:8128 | ✓ 1041ms | 否 | ✓ 993ms | ✓ 1259ms | ✓ 1247ms | http |
| 194.59.204.87:9080 | ✓ 1553ms | 否 | ✓ 1039ms | ✓ 1986ms | ✓ 1785ms | http |
| 129.213.162.27:17777 | ✓ 481ms | 否 | ✓ 1376ms | 否 | ✓ 859ms | http |
| 85.9.195.140:1080 | ✓ 475ms | ✓ 1208ms | ✓ 980ms | ✓ 1524ms | 否 | http |
| 103.67.46.225:3125 | ✓ 1987ms | 否 | ✓ 1873ms | ✓ 1763ms | ✓ 1637ms | http |
| 103.39.51.190:8080 | 否 | 否 | ✓ 1746ms | ✓ 1690ms | ✓ 1544ms | http |
| 88.80.150.82:8080 | ✓ 914ms | 否 | ✓ 907ms | 否 | ✓ 1441ms | https |
| 113.255.59.226:8080 | ✓ 1337ms | 否 | 否 | ✓ 1120ms | ✓ 1143ms | http |

<!-- END PROXY TABLE -->

## 🤝 参与贡献

本项目由社区驱动，欢迎任何形式的贡献。最简单的参与方式就是添加新的代理数据源。

请先阅读 **[贡献指南](CONTRIBUTING.md)** 了解如何开始。

## 🙏 支持本项目

如果您觉得本项目有帮助，欢迎给予支持，让更多人看到并参与贡献。

-   在 GitHub 上 **给本仓库加星** ⭐️
-   **分享**给朋友和同事

## ⚠️ 免责声明

-   本仓库中的代理均来自公开来源，不保证其速度、安全性或可用性。
-   使用这些代理的风险由您自行承担。
-   本仓库维护者不对任何滥用行为负责。请勿将代理用于非法用途。

## 📝 许可证

本仓库采用 MIT 许可证发布。详见 [LICENSE](LICENSE)。

## Stars
[![Star History Chart](https://api.star-history.com/svg?repos=feitianyul/free-proxy-list&type=Date)](https://star-history.com/#feitianyul/free-proxy-list&Date)

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

最后更新：2026-03-01 16:27:54 UTC（2026-03-02 00:27:54 UTC+8）

**代理总数：66**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 66 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 66 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 103.113.70.189:1081 | 否 | ✓ 1154ms | 否 | ✓ 1087ms | ✓ 879ms | http |
| 141.11.210.35:1080 | ✓ 1889ms | 否 | ✓ 1044ms | ✓ 1237ms | ✓ 918ms | http |
| 205.209.118.30:3138 | ✓ 927ms | ✓ 1982ms | ✓ 926ms | ✓ 1954ms | 否 | http |
| 104.238.30.45:59741 | ✓ 1613ms | 否 | ✓ 1807ms | 否 | ✓ 1903ms | http |
| 104.238.30.50:59741 | ✓ 1610ms | 否 | ✓ 1779ms | 否 | ✓ 1899ms | http |
| 217.76.245.80:999 | 否 | ✓ 1371ms | ✓ 1129ms | ✓ 1303ms | ✓ 1264ms | http |
| 104.238.30.91:63900 | ✓ 1668ms | 否 | ✓ 1747ms | 否 | ✓ 1899ms | http |
| 190.9.109.202:999 | ✓ 823ms | 否 | ✓ 1146ms | ✓ 1312ms | ✓ 1258ms | http |
| 120.92.212.16:8890 | 否 | ✓ 1460ms | ✓ 1382ms | ✓ 1677ms | ✓ 1160ms | http |
| 120.92.212.16:7890 | ✓ 1791ms | ✓ 1471ms | 否 | ✓ 1743ms | 否 | http |
| 104.238.30.86:63900 | ✓ 1870ms | 否 | ✓ 1903ms | 否 | ✓ 1967ms | http |
| 81.70.169.194:80 | ✓ 1160ms | ✓ 1568ms | ✓ 1245ms | ✓ 1544ms | 否 | http |
| 103.104.99.29:80 | 否 | 否 | ✓ 1946ms | ✓ 1802ms | ✓ 1763ms | http |
| 103.104.99.89:80 | 否 | 否 | ✓ 1953ms | ✓ 1815ms | ✓ 1844ms | http |
| 62.113.119.14:8080 | ✓ 1058ms | ✓ 1495ms | ✓ 1169ms | ✓ 1455ms | ✓ 1091ms | http |
| 101.43.255.96:80 | ✓ 1613ms | ✓ 1559ms | ✓ 1751ms | ✓ 1464ms | ✓ 1260ms | http |
| 95.85.252.153:21064 | ✓ 435ms | ✓ 1751ms | ✓ 1606ms | 否 | 否 | http |
| 103.84.95.54:7890 | ✓ 873ms | 否 | 否 | ✓ 1438ms | ✓ 1022ms | http |
| 196.70.95.87:3128 | ✓ 829ms | 否 | 否 | ✓ 1721ms | ✓ 1062ms | http |
| 104.238.30.40:59741 | ✓ 1617ms | 否 | ✓ 1648ms | 否 | ✓ 1903ms | http |
| 211.95.152.50:45046 | ✓ 1249ms | ✓ 1547ms | ✓ 1422ms | 否 | ✓ 1231ms | http |
| 200.125.171.254:999 | ✓ 769ms | ✓ 1256ms | ✓ 1388ms | ✓ 1386ms | ✓ 1097ms | http |
| 148.135.85.87:1080 | ✓ 1378ms | ✓ 1717ms | ✓ 1961ms | ✓ 1584ms | ✓ 1899ms | http |
| 2.56.178.131:443 | 否 | 否 | ✓ 1428ms | ✓ 1963ms | ✓ 1655ms | http |
| 168.235.110.63:3128 | ✓ 351ms | 否 | ✓ 1125ms | ✓ 1339ms | 否 | http |
| 45.22.209.157:8888 | 否 | 否 | ✓ 1430ms | ✓ 1431ms | ✓ 1858ms | http |
| 165.227.5.10:8888 | 否 | 否 | ✓ 1284ms | ✓ 1999ms | ✓ 1825ms | http |
| 59.46.216.131:30001 | ✓ 1179ms | 否 | ✓ 1594ms | 否 | ✓ 1247ms | http |
| 35.225.22.61:80 | ✓ 591ms | 否 | ✓ 1407ms | ✓ 1078ms | ✓ 1122ms | http |
| 36.155.100.217:8080 | ✓ 1794ms | ✓ 1560ms | ✓ 1323ms | 否 | ✓ 1348ms | http |
| 222.228.171.92:8080 | ✓ 1622ms | 否 | ✓ 1637ms | ✓ 1367ms | ✓ 1615ms | http |
| 74.208.234.198:443 | ✓ 610ms | 否 | ✓ 1366ms | ✓ 1567ms | ✓ 961ms | http |
| 103.215.36.88:15088 | 否 | ✓ 1426ms | ✓ 1426ms | ✓ 1574ms | 否 | http |
| 45.125.67.37:8443 | ✓ 1106ms | 否 | 否 | ✓ 1461ms | ✓ 1064ms | http |
| 115.231.181.40:8128 | ✓ 1533ms | ✓ 1363ms | 否 | 否 | ✓ 1100ms | http |
| 104.238.30.39:59741 | ✓ 1613ms | 否 | ✓ 1711ms | 否 | ✓ 1903ms | http |
| 45.140.147.155:1082 | ✓ 455ms | 否 | ✓ 1654ms | ✓ 1376ms | ✓ 1182ms | http |
| 34.101.184.164:3128 | ✓ 1313ms | 否 | ✓ 1843ms | ✓ 1893ms | ✓ 1239ms | http |
| 45.140.147.155:1081 | ✓ 454ms | ✓ 1417ms | 否 | ✓ 1635ms | ✓ 1154ms | http |
| 1.225.116.115:1080 | ✓ 1562ms | 否 | 否 | ✓ 1314ms | ✓ 930ms | http |
| 36.147.78.166:80 | 否 | ✓ 1893ms | 否 | ✓ 1929ms | ✓ 1619ms | http |
| 198.244.151.77:8888 | ✓ 1691ms | 否 | ✓ 1411ms | 否 | ✓ 1723ms | http |
| 5.75.201.136:1080 | ✓ 834ms | ✓ 1411ms | 否 | ✓ 1495ms | ✓ 1756ms | http |
| 101.32.244.83:8080 | ✓ 1210ms | 否 | ✓ 1183ms | ✓ 1699ms | ✓ 1503ms | http |
| 121.43.196.213:8222 | ✓ 1149ms | ✓ 1282ms | ✓ 1060ms | ✓ 1392ms | ✓ 1031ms | http |
| 121.43.196.210:8222 | ✓ 1142ms | ✓ 1340ms | ✓ 1017ms | ✓ 1340ms | ✓ 1087ms | http |
| 114.55.226.123:10086 | ✓ 1572ms | ✓ 1933ms | ✓ 1168ms | ✓ 1532ms | ✓ 1196ms | http |
| 35.234.17.221:8080 | ✓ 1168ms | 否 | 否 | ✓ 1223ms | ✓ 1264ms | http |
| 85.208.108.43:2094 | ✓ 1126ms | 否 | ✓ 1264ms | 否 | ✓ 753ms | http |
| 104.238.30.37:59741 | ✓ 1633ms | 否 | ✓ 1614ms | 否 | ✓ 1907ms | http |
| 89.169.168.25:3128 | ✓ 864ms | ✓ 1867ms | ✓ 839ms | 否 | 否 | http |
| 14.56.107.244:3128 | 否 | 否 | ✓ 1174ms | ✓ 1247ms | ✓ 1294ms | http |
| 167.160.184.231:6005 | ✓ 854ms | ✓ 1602ms | ✓ 1192ms | 否 | ✓ 1181ms | http |
| 125.129.39.95:3128 | ✓ 1283ms | 否 | ✓ 1108ms | ✓ 1535ms | ✓ 1360ms | http |
| 183.96.6.112:3128 | ✓ 1283ms | ✓ 1587ms | ✓ 1524ms | ✓ 1535ms | ✓ 1355ms | http |
| 62.234.206.73:3128 | ✓ 1193ms | 否 | ✓ 1085ms | ✓ 1529ms | ✓ 1441ms | http |
| 45.136.198.40:3128 | ✓ 754ms | 否 | ✓ 1505ms | 否 | ✓ 1710ms | http |
| 104.238.30.38:59741 | ✓ 1572ms | 否 | ✓ 1643ms | 否 | ✓ 1936ms | http |
| 172.212.68.37:3128 | ✓ 535ms | 否 | ✓ 695ms | ✓ 986ms | ✓ 840ms | http |
| 183.128.208.68:7890 | ✓ 1485ms | 否 | ✓ 1073ms | 否 | ✓ 1190ms | http |
| 217.77.102.18:3128 | ✓ 1619ms | 否 | ✓ 1168ms | 否 | ✓ 1829ms | http |
| 85.198.84.77:10808 | ✓ 1640ms | 否 | ✓ 1813ms | 否 | ✓ 1534ms | http |
| 61.52.131.172:8443 | ✓ 1029ms | ✓ 1407ms | ✓ 1132ms | ✓ 1370ms | 否 | http |
| 103.39.51.190:8080 | ✓ 1967ms | 否 | ✓ 1462ms | ✓ 1547ms | ✓ 1622ms | http |
| 103.215.36.88:15968 | ✓ 1255ms | ✓ 1480ms | ✓ 1161ms | 否 | 否 | http |
| 180.127.149.247:1080 | ✓ 1349ms | ✓ 1491ms | 否 | 否 | ✓ 1129ms | http |

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

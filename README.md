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

最后更新：2026-03-03 06:52:37 UTC（2026-03-03 14:52:37 UTC+8）

**代理总数：65**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 65 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 65 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 3.213.157.4:3128 | ✓ 350ms | 否 | ✓ 131ms | 否 | ✓ 921ms | http |
| 120.238.159.228:22222 | ✓ 1069ms | ✓ 1345ms | ✓ 961ms | ✓ 1265ms | ✓ 961ms | http |
| 183.249.5.109:22222 | ✓ 1087ms | ✓ 1477ms | 否 | ✓ 1244ms | ✓ 830ms | http |
| 45.88.0.113:3128 | ✓ 531ms | 否 | 否 | ✓ 1285ms | ✓ 1336ms | http |
| 166.0.192.117:8888 | ✓ 1481ms | 否 | ✓ 1663ms | ✓ 1823ms | ✓ 1253ms | http |
| 120.232.242.119:22222 | ✓ 1276ms | ✓ 1330ms | 否 | ✓ 1620ms | 否 | http |
| 120.92.212.16:8890 | ✓ 1070ms | 否 | 否 | ✓ 1571ms | ✓ 1989ms | http |
| 186.148.180.46:999 | ✓ 1283ms | 否 | ✓ 1492ms | 否 | ✓ 1365ms | http |
| 125.128.12.14:3128 | 否 | ✓ 1593ms | ✓ 700ms | ✓ 1177ms | ✓ 855ms | http |
| 205.209.118.30:3138 | ✓ 270ms | 否 | ✓ 132ms | ✓ 1114ms | ✓ 838ms | http |
| 61.72.221.94:3128 | ✓ 1383ms | 否 | ✓ 1365ms | 否 | ✓ 1269ms | http |
| 61.72.110.54:3128 | 否 | 否 | ✓ 890ms | ✓ 1190ms | ✓ 1946ms | http |
| 103.131.19.42:8181 | ✓ 1560ms | 否 | ✓ 1705ms | 否 | ✓ 1823ms | http |
| 35.225.22.61:80 | ✓ 285ms | 否 | ✓ 424ms | ✓ 813ms | ✓ 752ms | http |
| 183.249.5.117:22222 | ✓ 787ms | ✓ 996ms | ✓ 875ms | 否 | ✓ 823ms | http |
| 103.84.95.54:7890 | 否 | 否 | ✓ 904ms | ✓ 1189ms | ✓ 1456ms | http |
| 120.92.212.16:7890 | ✓ 1106ms | 否 | 否 | ✓ 1615ms | ✓ 1085ms | http |
| 160.238.65.3:3128 | ✓ 630ms | 否 | ✓ 1935ms | ✓ 1945ms | ✓ 1463ms | http |
| 160.238.65.2:3128 | ✓ 1801ms | 否 | ✓ 1830ms | ✓ 1697ms | ✓ 1014ms | http |
| 120.240.35.161:22222 | 否 | ✓ 1701ms | ✓ 1270ms | 否 | ✓ 1203ms | http |
| 160.238.65.4:3128 | ✓ 792ms | 否 | ✓ 772ms | 否 | ✓ 1330ms | http |
| 81.70.169.194:80 | ✓ 1806ms | ✓ 1514ms | ✓ 1107ms | ✓ 1351ms | ✓ 1360ms | http |
| 101.43.255.96:80 | ✓ 1197ms | ✓ 1378ms | ✓ 1355ms | ✓ 1637ms | 否 | http |
| 160.238.65.8:3128 | ✓ 523ms | ✓ 1519ms | ✓ 448ms | 否 | 否 | http |
| 117.159.239.52:22222 | ✓ 955ms | ✓ 1207ms | ✓ 913ms | ✓ 1246ms | ✓ 964ms | http |
| 160.250.5.22:1 | ✓ 1007ms | 否 | ✓ 1431ms | ✓ 1752ms | ✓ 1246ms | http |
| 160.250.4.245:1 | ✓ 1066ms | 否 | ✓ 1984ms | ✓ 1562ms | ✓ 1120ms | http |
| 35.234.17.221:8080 | ✓ 1298ms | 否 | 否 | ✓ 1069ms | ✓ 949ms | http |
| 120.240.35.174:22222 | ✓ 1425ms | ✓ 1744ms | 否 | ✓ 1527ms | ✓ 1144ms | http |
| 221.127.195.224:8888 | ✓ 1424ms | 否 | ✓ 1299ms | ✓ 1459ms | ✓ 1305ms | http |
| 121.230.9.4:1080 | ✓ 1702ms | ✓ 1610ms | ✓ 1367ms | 否 | 否 | http |
| 180.127.149.247:1080 | ✓ 1907ms | ✓ 1298ms | ✓ 1128ms | 否 | 否 | http |
| 5.75.196.26:40000 | ✓ 473ms | ✓ 1810ms | ✓ 1135ms | ✓ 1496ms | ✓ 1373ms | http |
| 138.124.53.25:7443 | 否 | 否 | ✓ 1782ms | ✓ 1871ms | ✓ 1825ms | http |
| 91.193.240.157:9877 | ✓ 1470ms | 否 | ✓ 1695ms | 否 | ✓ 1591ms | http |
| 5.102.109.41:999 | ✓ 508ms | ✓ 1401ms | ✓ 1470ms | ✓ 1331ms | 否 | http |
| 45.136.198.40:3128 | ✓ 766ms | 否 | ✓ 1569ms | 否 | ✓ 1929ms | http |
| 121.230.8.154:1080 | 否 | 否 | ✓ 1598ms | ✓ 1865ms | ✓ 1777ms | http |
| 183.249.5.214:22222 | 否 | ✓ 1259ms | 否 | ✓ 1082ms | ✓ 1102ms | http |
| 45.77.249.199:1236 | 否 | 否 | ✓ 1869ms | ✓ 1176ms | ✓ 1151ms | http |
| 120.240.35.160:22222 | ✓ 1221ms | ✓ 1722ms | ✓ 1127ms | ✓ 1462ms | ✓ 1160ms | http |
| 47.101.149.27:9010 | 否 | ✓ 1431ms | ✓ 1466ms | ✓ 1619ms | ✓ 1309ms | http |
| 142.171.131.38:7890 | ✓ 1957ms | ✓ 1614ms | ✓ 1981ms | ✓ 925ms | 否 | http |
| 121.230.9.5:1080 | ✓ 1908ms | ✓ 1591ms | 否 | ✓ 1866ms | 否 | http |
| 95.85.252.153:21064 | ✓ 446ms | 否 | ✓ 1152ms | ✓ 1868ms | 否 | http |
| 120.240.29.51:22222 | 否 | ✓ 1518ms | ✓ 1045ms | ✓ 1285ms | ✓ 1024ms | http |
| 2.56.178.131:443 | 否 | 否 | ✓ 1722ms | ✓ 1759ms | ✓ 1794ms | http |
| 45.88.0.117:3128 | ✓ 893ms | 否 | ✓ 1546ms | 否 | ✓ 1079ms | http |
| 94.177.131.12:3128 | ✓ 1151ms | 否 | ✓ 636ms | ✓ 948ms | ✓ 1030ms | http |
| 46.249.103.192:443 | ✓ 620ms | 否 | ✓ 1920ms | 否 | ✓ 1452ms | http |
| 129.212.226.87:443 | ✓ 1771ms | 否 | ✓ 1102ms | ✓ 1276ms | 否 | http |
| 209.38.51.97:3128 | ✓ 366ms | 否 | ✓ 815ms | 否 | ✓ 987ms | http |
| 117.159.239.51:22222 | ✓ 975ms | ✓ 1203ms | ✓ 939ms | ✓ 1224ms | ✓ 933ms | http |
| 120.240.35.177:22222 | ✓ 1109ms | 否 | 否 | ✓ 1531ms | ✓ 1115ms | http |
| 120.198.141.75:22222 | ✓ 1155ms | ✓ 1399ms | ✓ 987ms | ✓ 1388ms | ✓ 1046ms | http |
| 121.128.121.54:3128 | 否 | 否 | ✓ 1235ms | ✓ 1375ms | ✓ 1025ms | http |
| 154.90.48.209:9090 | ✓ 1661ms | 否 | ✓ 1708ms | ✓ 1662ms | 否 | http |
| 103.166.185.54:3128 | ✓ 1863ms | 否 | ✓ 1374ms | ✓ 1433ms | ✓ 1151ms | http |
| 103.39.51.190:8080 | ✓ 1610ms | 否 | ✓ 1378ms | ✓ 1537ms | ✓ 1597ms | http |
| 192.71.213.85:9091 | ✓ 737ms | 否 | ✓ 834ms | ✓ 1831ms | 否 | http |
| 47.106.73.57:8118 | 否 | 否 | ✓ 1605ms | ✓ 1977ms | ✓ 1054ms | http |
| 183.249.5.111:22222 | ✓ 971ms | ✓ 1326ms | ✓ 829ms | ✓ 1110ms | ✓ 1006ms | http |
| 113.59.32.141:22222 | 否 | ✓ 1576ms | ✓ 1204ms | 否 | ✓ 1167ms | http |
| 115.231.181.40:8128 | 否 | ✓ 1593ms | 否 | ✓ 1283ms | ✓ 1308ms | http |
| 172.212.68.37:3128 | ✓ 300ms | ✓ 1451ms | ✓ 739ms | ✓ 1811ms | ✓ 760ms | http |

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

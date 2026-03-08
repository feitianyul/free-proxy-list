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

最后更新：2026-03-08 10:29:23 UTC（2026-03-08 18:29:23 UTC+8）

**代理总数：53**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 52 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 53 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 181.78.79.155:999 | ✓ 1633ms | 否 | 否 | ✓ 1664ms | ✓ 1689ms | http |
| 205.209.118.30:3138 | ✓ 1647ms | ✓ 968ms | ✓ 77ms | 否 | 否 | http |
| 194.213.18.200:443 | ✓ 1188ms | 否 | 否 | ✓ 1864ms | ✓ 1582ms | http |
| 103.84.95.54:7890 | ✓ 870ms | 否 | ✓ 948ms | ✓ 1843ms | ✓ 1523ms | http |
| 185.243.218.43:49153 | ✓ 819ms | 否 | ✓ 1701ms | ✓ 1903ms | ✓ 1791ms | http |
| 120.92.212.16:8890 | ✓ 1108ms | ✓ 1397ms | 否 | ✓ 1670ms | ✓ 1153ms | http |
| 120.92.212.16:7890 | ✓ 1396ms | 否 | 否 | ✓ 1723ms | ✓ 1425ms | http |
| 120.240.35.176:22222 | ✓ 1855ms | ✓ 1439ms | ✓ 1144ms | ✓ 1448ms | 否 | http |
| 202.155.12.161:443 | ✓ 1628ms | 否 | ✓ 1359ms | ✓ 1397ms | ✓ 1106ms | http |
| 46.249.103.192:443 | ✓ 843ms | 否 | ✓ 1518ms | ✓ 1449ms | 否 | http |
| 1.231.81.166:3128 | 否 | 否 | ✓ 1379ms | ✓ 1202ms | ✓ 1187ms | http |
| 81.70.169.194:80 | ✓ 1079ms | 否 | 否 | ✓ 1496ms | ✓ 1166ms | http |
| 101.43.255.96:80 | 否 | ✓ 1468ms | 否 | ✓ 1440ms | ✓ 1220ms | http |
| 120.240.35.161:22222 | 否 | ✓ 1740ms | 否 | ✓ 1417ms | ✓ 1107ms | http |
| 47.77.193.180:1080 | 否 | ✓ 1042ms | ✓ 791ms | ✓ 974ms | ✓ 729ms | http |
| 88.80.150.82:8080 | ✓ 1298ms | ✓ 1911ms | 否 | 否 | ✓ 1548ms | https |
| 103.82.23.118:5253 | ✓ 1338ms | 否 | ✓ 1533ms | 否 | ✓ 1437ms | http |
| 222.184.48.252:22222 | ✓ 1899ms | 否 | 否 | ✓ 1552ms | ✓ 1113ms | http |
| 103.215.36.88:18653 | ✓ 1520ms | 否 | ✓ 1255ms | ✓ 1924ms | ✓ 1646ms | http |
| 203.205.49.2:10017 | ✓ 1417ms | 否 | ✓ 1679ms | ✓ 1449ms | ✓ 1668ms | http |
| 39.104.201.40:7890 | 否 | ✓ 1423ms | ✓ 1568ms | ✓ 1204ms | 否 | http |
| 113.59.32.162:22222 | ✓ 1330ms | ✓ 1593ms | ✓ 1339ms | ✓ 1539ms | ✓ 1231ms | http |
| 101.47.73.135:3128 | ✓ 1640ms | 否 | 否 | ✓ 1297ms | ✓ 1371ms | http |
| 120.238.159.227:22222 | ✓ 1238ms | ✓ 1507ms | ✓ 1035ms | ✓ 1461ms | ✓ 1049ms | http |
| 120.240.35.173:22222 | ✓ 1104ms | ✓ 1478ms | ✓ 1214ms | ✓ 1418ms | ✓ 1097ms | http |
| 120.198.141.75:22222 | 否 | ✓ 1460ms | ✓ 1138ms | 否 | ✓ 1091ms | http |
| 183.249.5.213:22222 | ✓ 939ms | ✓ 1122ms | 否 | ✓ 1151ms | ✓ 1091ms | http |
| 120.240.29.168:22222 | ✓ 1499ms | 否 | ✓ 1300ms | 否 | ✓ 1088ms | http |
| 106.14.205.114:483 | 否 | ✓ 1890ms | ✓ 1380ms | ✓ 1394ms | ✓ 1046ms | http |
| 42.115.72.27:2033 | 否 | 否 | ✓ 1702ms | ✓ 1944ms | ✓ 1935ms | http |
| 45.140.147.82:1081 | ✓ 1354ms | 否 | ✓ 1365ms | 否 | ✓ 1116ms | http |
| 165.227.5.10:8888 | ✓ 588ms | 否 | 否 | ✓ 1022ms | ✓ 1772ms | http |
| 180.130.80.196:9003 | ✓ 1252ms | ✓ 1391ms | ✓ 1414ms | 否 | 否 | http |
| 116.80.82.226:3172 | ✓ 1635ms | 否 | ✓ 1724ms | 否 | ✓ 1786ms | http |
| 172.212.68.37:3128 | ✓ 240ms | ✓ 1415ms | ✓ 942ms | ✓ 1339ms | ✓ 1002ms | http |
| 159.89.31.62:8080 | ✓ 872ms | ✓ 1643ms | ✓ 1507ms | ✓ 1825ms | 否 | http |
| 222.184.48.241:22222 | 否 | 否 | ✓ 1061ms | ✓ 1741ms | ✓ 1387ms | http |
| 45.136.198.40:3128 | ✓ 667ms | 否 | ✓ 1998ms | ✓ 1897ms | ✓ 1397ms | http |
| 103.215.36.88:17133 | ✓ 1315ms | 否 | ✓ 1592ms | ✓ 1482ms | ✓ 1243ms | http |
| 103.215.36.88:19787 | ✓ 1141ms | ✓ 1496ms | 否 | 否 | ✓ 1216ms | http |
| 103.215.36.88:17083 | ✓ 1809ms | ✓ 1709ms | 否 | ✓ 1518ms | ✓ 1606ms | http |
| 117.159.239.49:22222 | ✓ 1068ms | ✓ 1354ms | ✓ 967ms | ✓ 1268ms | ✓ 1032ms | http |
| 120.198.141.79:22222 | ✓ 1126ms | ✓ 1385ms | ✓ 1074ms | ✓ 1514ms | 否 | http |
| 120.240.35.177:22222 | ✓ 1046ms | ✓ 1389ms | ✓ 1143ms | ✓ 1402ms | ✓ 1181ms | http |
| 59.46.216.131:30001 | 否 | ✓ 1491ms | ✓ 1259ms | 否 | ✓ 1187ms | http |
| 168.235.110.63:3128 | ✓ 1746ms | 否 | ✓ 1200ms | ✓ 1416ms | 否 | http |
| 45.186.6.104:3128 | ✓ 1678ms | ✓ 1904ms | ✓ 1741ms | 否 | 否 | http |
| 8.137.112.117:3128 | ✓ 975ms | ✓ 1424ms | ✓ 977ms | ✓ 1451ms | ✓ 1231ms | http |
| 113.177.131.2:3128 | ✓ 1813ms | 否 | ✓ 1124ms | ✓ 1997ms | ✓ 1197ms | http |
| 115.231.181.40:8128 | ✓ 1526ms | ✓ 1821ms | ✓ 1822ms | ✓ 1707ms | 否 | http |
| 85.9.195.140:1080 | ✓ 335ms | 否 | ✓ 1949ms | 否 | ✓ 1801ms | http |
| 103.39.51.190:8080 | ✓ 1987ms | 否 | ✓ 1458ms | 否 | ✓ 1711ms | http |
| 120.232.242.119:22222 | ✓ 1054ms | 否 | ✓ 1232ms | ✓ 1344ms | 否 | http |

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

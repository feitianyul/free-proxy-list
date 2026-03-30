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

最后更新：2026-03-30 11:58:19 UTC（2026-03-30 19:58:19 UTC+8）

**代理总数：59**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 59 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 59 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 208.87.243.199:7878 | ✓ 883ms | ✓ 896ms | ✓ 1004ms | ✓ 1085ms | ✓ 1925ms | http |
| 147.161.239.240:8800 | ✓ 692ms | ✓ 1490ms | ✓ 687ms | ✓ 1564ms | ✓ 1313ms | http |
| 39.185.46.193:5911 | ✓ 904ms | ✓ 1059ms | ✓ 915ms | ✓ 1160ms | ✓ 890ms | http |
| 103.84.95.54:7890 | ✓ 995ms | 否 | ✓ 989ms | ✓ 1425ms | ✓ 1355ms | http |
| 147.161.210.140:8800 | ✓ 1747ms | 否 | ✓ 1122ms | ✓ 1390ms | ✓ 1227ms | http |
| 1.231.81.166:3128 | ✓ 1780ms | ✓ 1511ms | ✓ 1847ms | ✓ 1439ms | ✓ 1208ms | http |
| 167.103.115.102:8800 | ✓ 1604ms | 否 | ✓ 1146ms | ✓ 1513ms | ✓ 1573ms | http |
| 113.160.132.26:8080 | ✓ 1913ms | ✓ 1716ms | ✓ 1533ms | 否 | 否 | http |
| 115.231.181.40:8128 | ✓ 1171ms | 否 | ✓ 1114ms | ✓ 1418ms | ✓ 1085ms | http |
| 42.96.16.158:1311 | ✓ 1642ms | 否 | 否 | ✓ 1634ms | ✓ 1111ms | http |
| 43.99.54.236:5555 | ✓ 907ms | ✓ 1192ms | ✓ 837ms | ✓ 1051ms | ✓ 890ms | http |
| 183.249.5.105:22222 | ✓ 1018ms | ✓ 1381ms | ✓ 1091ms | ✓ 1317ms | ✓ 1066ms | http |
| 45.136.198.40:3128 | ✓ 1324ms | ✓ 1806ms | ✓ 1495ms | ✓ 1813ms | ✓ 1844ms | http |
| 167.103.34.108:8800 | ✓ 1261ms | 否 | ✓ 1221ms | ✓ 1432ms | ✓ 1430ms | http |
| 120.92.212.16:8890 | ✓ 1226ms | 否 | 否 | ✓ 1468ms | ✓ 1385ms | http |
| 222.184.48.251:22222 | ✓ 1036ms | ✓ 1358ms | ✓ 1194ms | ✓ 1479ms | ✓ 1816ms | http |
| 35.225.22.61:80 | 否 | 否 | ✓ 983ms | ✓ 990ms | ✓ 847ms | http |
| 95.213.217.168:52004 | ✓ 513ms | ✓ 1507ms | ✓ 838ms | ✓ 1511ms | ✓ 1082ms | http |
| 45.12.151.226:2829 | ✓ 550ms | ✓ 1585ms | 否 | ✓ 1567ms | ✓ 1212ms | http |
| 167.103.144.127:8800 | ✓ 1448ms | 否 | ✓ 1378ms | ✓ 1601ms | ✓ 1332ms | http |
| 103.166.185.54:3128 | ✓ 1628ms | ✓ 1829ms | ✓ 1363ms | ✓ 1813ms | ✓ 1154ms | http |
| 167.103.31.122:8800 | ✓ 1604ms | 否 | ✓ 1285ms | ✓ 1930ms | ✓ 1381ms | http |
| 190.12.150.244:999 | 否 | ✓ 1838ms | ✓ 1150ms | ✓ 1631ms | ✓ 1643ms | http |
| 85.208.108.43:2094 | 否 | 否 | ✓ 1360ms | ✓ 926ms | ✓ 719ms | http |
| 183.249.5.117:22222 | ✓ 1216ms | ✓ 1400ms | ✓ 848ms | ✓ 1214ms | ✓ 968ms | http |
| 59.46.216.131:30001 | ✓ 1162ms | 否 | 否 | ✓ 1646ms | ✓ 1325ms | http |
| 117.159.239.49:22222 | ✓ 989ms | ✓ 1418ms | ✓ 1005ms | ✓ 1334ms | ✓ 1062ms | http |
| 222.184.48.242:22222 | 否 | 否 | ✓ 1470ms | ✓ 1916ms | ✓ 1098ms | http |
| 209.126.84.232:8888 | 否 | ✓ 1277ms | ✓ 1387ms | ✓ 1692ms | ✓ 980ms | http |
| 101.43.127.100:8877 | 否 | ✓ 1328ms | ✓ 1110ms | ✓ 1429ms | 否 | http |
| 45.129.141.143:3128 | ✓ 595ms | ✓ 1630ms | ✓ 1479ms | ✓ 1942ms | ✓ 1600ms | http |
| 86.53.183.16:1080 | ✓ 858ms | ✓ 1803ms | ✓ 1642ms | 否 | 否 | http |
| 183.249.5.110:22222 | ✓ 1691ms | 否 | ✓ 1608ms | 否 | ✓ 1341ms | http |
| 222.184.48.252:22222 | ✓ 1961ms | ✓ 1412ms | 否 | 否 | ✓ 1133ms | http |
| 103.154.214.50:3128 | ✓ 1469ms | 否 | ✓ 1722ms | ✓ 1397ms | ✓ 1071ms | http |
| 120.92.212.16:7890 | ✓ 1947ms | ✓ 1698ms | ✓ 1549ms | ✓ 1976ms | ✓ 1601ms | http |
| 177.234.217.88:999 | 否 | ✓ 1926ms | ✓ 1738ms | ✓ 1674ms | ✓ 1353ms | http |
| 104.247.51.76:3128 | ✓ 1662ms | 否 | ✓ 1058ms | ✓ 1024ms | ✓ 815ms | http |
| 45.140.147.82:1082 | ✓ 1419ms | 否 | ✓ 459ms | ✓ 1084ms | ✓ 1085ms | http |
| 45.140.147.82:1081 | ✓ 555ms | ✓ 1828ms | ✓ 678ms | 否 | ✓ 969ms | http |
| 185.191.236.162:3128 | ✓ 1630ms | ✓ 1933ms | ✓ 1885ms | 否 | 否 | http |
| 101.32.244.83:8080 | ✓ 1669ms | ✓ 1946ms | ✓ 1179ms | ✓ 1711ms | ✓ 1538ms | http |
| 121.43.196.210:8222 | ✓ 1143ms | ✓ 1255ms | ✓ 1127ms | ✓ 1385ms | ✓ 1090ms | http |
| 121.43.196.213:8222 | ✓ 1186ms | ✓ 1386ms | ✓ 1018ms | ✓ 1352ms | ✓ 1070ms | http |
| 114.55.226.123:10086 | ✓ 1881ms | ✓ 1672ms | ✓ 1202ms | ✓ 1497ms | ✓ 1216ms | http |
| 91.233.223.147:3128 | ✓ 1101ms | 否 | ✓ 1136ms | 否 | ✓ 1494ms | http |
| 222.184.48.241:22222 | ✓ 1956ms | ✓ 1441ms | 否 | ✓ 1423ms | ✓ 1466ms | http |
| 47.74.226.8:5001 | ✓ 1197ms | ✓ 1821ms | ✓ 1207ms | 否 | 否 | http |
| 76.169.128.104:8080 | ✓ 671ms | 否 | ✓ 917ms | ✓ 1686ms | 否 | http |
| 34.101.184.164:3128 | ✓ 981ms | 否 | ✓ 1629ms | 否 | ✓ 1117ms | http |
| 103.39.51.190:8080 | ✓ 1956ms | 否 | 否 | ✓ 1840ms | ✓ 1591ms | http |
| 180.130.80.196:9003 | ✓ 1498ms | ✓ 1946ms | 否 | ✓ 1808ms | ✓ 1294ms | http |
| 5.102.109.41:999 | ✓ 1390ms | 否 | ✓ 1748ms | ✓ 1855ms | 否 | http |
| 183.249.5.111:22222 | ✓ 1021ms | ✓ 1178ms | 否 | ✓ 1181ms | ✓ 958ms | http |
| 114.237.77.231:1080 | ✓ 1499ms | ✓ 1348ms | ✓ 1166ms | 否 | ✓ 1045ms | http |
| 31.192.106.135:8010 | ✓ 1113ms | ✓ 1742ms | ✓ 1885ms | 否 | 否 | http |
| 121.42.162.62:1110 | ✓ 1788ms | ✓ 1604ms | ✓ 1738ms | 否 | ✓ 1699ms | http |
| 59.8.203.55:80 | 否 | ✓ 1442ms | ✓ 1468ms | ✓ 1222ms | ✓ 999ms | http |
| 217.76.245.80:999 | ✓ 405ms | 否 | ✓ 615ms | 否 | ✓ 1018ms | http |

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

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

最后更新：2026-03-07 12:17:45 UTC（2026-03-07 20:17:45 UTC+8）

**代理总数：51**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 51 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 51 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 205.209.118.30:3138 | ✓ 354ms | 否 | ✓ 1155ms | ✓ 1342ms | ✓ 1059ms | http |
| 103.84.95.54:7890 | ✓ 875ms | 否 | ✓ 1690ms | 否 | ✓ 1712ms | http |
| 217.76.245.80:999 | ✓ 1139ms | ✓ 1661ms | ✓ 1319ms | ✓ 1687ms | ✓ 1462ms | http |
| 46.249.103.192:443 | ✓ 992ms | 否 | ✓ 1697ms | ✓ 1933ms | 否 | http |
| 168.235.110.63:3128 | ✓ 685ms | 否 | ✓ 1140ms | ✓ 1354ms | 否 | http |
| 14.225.222.164:7890 | 否 | ✓ 1494ms | ✓ 1727ms | ✓ 1009ms | ✓ 1521ms | http |
| 167.172.69.123:8080 | ✓ 692ms | 否 | ✓ 985ms | ✓ 1045ms | ✓ 1509ms | http |
| 167.172.69.123:80 | ✓ 705ms | 否 | ✓ 1143ms | 否 | ✓ 970ms | http |
| 81.70.169.194:80 | ✓ 942ms | 否 | ✓ 1047ms | ✓ 1290ms | ✓ 1792ms | http |
| 101.43.255.96:80 | 否 | ✓ 1323ms | ✓ 1217ms | 否 | ✓ 908ms | http |
| 35.225.22.61:80 | 否 | 否 | ✓ 514ms | ✓ 1274ms | ✓ 869ms | http |
| 46.183.25.8:443 | ✓ 394ms | 否 | ✓ 373ms | ✓ 1016ms | ✓ 1301ms | http |
| 190.9.109.199:999 | ✓ 958ms | ✓ 1555ms | 否 | ✓ 1483ms | ✓ 1211ms | http |
| 106.14.203.63:3333 | ✓ 1047ms | ✓ 993ms | ✓ 803ms | ✓ 1085ms | ✓ 829ms | http |
| 103.215.36.88:16359 | ✓ 1122ms | ✓ 1402ms | ✓ 983ms | ✓ 1841ms | ✓ 907ms | http |
| 91.193.240.157:9877 | ✓ 1257ms | 否 | ✓ 1760ms | 否 | ✓ 1560ms | http |
| 183.81.91.186:2112 | ✓ 1928ms | 否 | ✓ 1405ms | ✓ 1485ms | 否 | http |
| 8.219.97.248:80 | ✓ 1273ms | 否 | ✓ 1364ms | ✓ 1937ms | 否 | http |
| 45.140.147.155:1081 | ✓ 653ms | ✓ 1612ms | ✓ 1435ms | ✓ 1476ms | ✓ 996ms | http |
| 120.92.212.16:8890 | 否 | ✓ 1159ms | ✓ 1637ms | ✓ 1191ms | 否 | http |
| 62.113.119.14:8080 | ✓ 787ms | ✓ 1792ms | ✓ 1209ms | ✓ 1680ms | ✓ 1232ms | http |
| 150.107.140.238:3128 | 否 | 否 | ✓ 1869ms | ✓ 1088ms | ✓ 999ms | http |
| 194.59.204.87:9080 | ✓ 1045ms | 否 | ✓ 937ms | 否 | ✓ 1457ms | http |
| 185.243.218.43:49153 | ✓ 1115ms | 否 | ✓ 1865ms | 否 | ✓ 1830ms | http |
| 101.32.244.83:8080 | ✓ 1425ms | 否 | ✓ 881ms | ✓ 1285ms | ✓ 994ms | http |
| 121.43.196.210:8222 | ✓ 908ms | ✓ 991ms | ✓ 884ms | ✓ 1058ms | ✓ 869ms | http |
| 121.43.196.213:8222 | ✓ 1010ms | ✓ 1061ms | ✓ 868ms | ✓ 1075ms | ✓ 879ms | http |
| 114.55.226.123:10086 | ✓ 974ms | ✓ 1601ms | ✓ 1072ms | ✓ 1206ms | ✓ 983ms | http |
| 103.139.138.194:3128 | ✓ 1715ms | 否 | ✓ 1449ms | ✓ 1295ms | ✓ 1071ms | http |
| 59.46.216.131:30001 | ✓ 1062ms | 否 | 否 | ✓ 1985ms | ✓ 1051ms | http |
| 85.9.195.140:1080 | ✓ 1313ms | 否 | ✓ 954ms | ✓ 1882ms | 否 | http |
| 103.215.36.88:17063 | ✓ 947ms | ✓ 1278ms | 否 | 否 | ✓ 1059ms | http |
| 103.215.36.88:18091 | ✓ 1670ms | ✓ 1364ms | ✓ 1012ms | 否 | ✓ 1880ms | http |
| 14.56.107.244:3128 | 否 | ✓ 1622ms | ✓ 1898ms | ✓ 1358ms | 否 | http |
| 125.128.12.144:3128 | ✓ 1650ms | ✓ 1812ms | ✓ 1582ms | ✓ 1258ms | ✓ 858ms | http |
| 1.231.81.166:3128 | ✓ 763ms | ✓ 1938ms | ✓ 865ms | ✓ 925ms | ✓ 656ms | http |
| 120.92.212.16:7890 | ✓ 973ms | ✓ 1203ms | 否 | ✓ 1237ms | 否 | http |
| 121.128.121.54:3128 | 否 | ✓ 980ms | ✓ 760ms | ✓ 1061ms | ✓ 1152ms | http |
| 91.233.223.147:3128 | ✓ 994ms | 否 | ✓ 1058ms | 否 | ✓ 1720ms | http |
| 202.129.206.239:3128 | ✓ 1833ms | 否 | ✓ 1574ms | ✓ 1549ms | ✓ 1456ms | http |
| 103.39.51.190:8080 | ✓ 1576ms | 否 | ✓ 1725ms | ✓ 1637ms | ✓ 1395ms | http |
| 188.132.141.249:443 | ✓ 1934ms | 否 | ✓ 1815ms | 否 | ✓ 1750ms | http |
| 16.78.119.130:443 | 否 | ✓ 1929ms | ✓ 1947ms | ✓ 1763ms | 否 | http |
| 178.236.245.17:3128 | ✓ 1120ms | 否 | ✓ 1616ms | 否 | ✓ 1513ms | http |
| 178.236.245.59:3128 | ✓ 1461ms | 否 | ✓ 1285ms | 否 | ✓ 1848ms | http |
| 115.231.181.40:8128 | ✓ 852ms | ✓ 1128ms | ✓ 914ms | ✓ 1113ms | 否 | http |
| 200.125.171.254:999 | ✓ 1685ms | 否 | ✓ 1311ms | ✓ 1459ms | ✓ 1321ms | http |
| 8.137.112.117:3128 | ✓ 923ms | ✓ 1252ms | ✓ 1888ms | 否 | 否 | http |
| 121.230.8.80:1080 | ✓ 1274ms | ✓ 1587ms | ✓ 1171ms | ✓ 1856ms | 否 | http |
| 45.136.198.40:3128 | 否 | 否 | ✓ 836ms | ✓ 1760ms | ✓ 1364ms | http |
| 45.140.147.82:1081 | 否 | ✓ 1759ms | ✓ 1660ms | ✓ 1894ms | 否 | http |

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
